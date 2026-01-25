package typo

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/invopop/jsonschema"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"

	"github.com/deigmata-paideias/typo/internal/repository"
	"github.com/deigmata-paideias/typo/internal/scanner"
	"github.com/deigmata-paideias/typo/internal/types"
	"github.com/deigmata-paideias/typo/internal/utils"
)

type ITypo interface {
	// Typo 返回修正的命令候选列表和原始命令
	Typo() (string, []types.MatchResult, error)
}

type LocalTypo struct {
	repo repository.IRepository
	hs   scanner.IScanner
}

func NewLocalTypo(repo repository.IRepository, hs scanner.IScanner) ITypo {

	return &LocalTypo{
		repo,
		hs,
	}
}

// Typo 返回修正的命令候选列表和原始命令
// first return: original command
// second return: candidate commands with scores
func (t *LocalTypo) Typo() (string, []types.MatchResult, error) {

	command, err := t.hs.Scan()
	if err != nil {
		return "", nil, err
	}

	// 分割命令
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return command, nil, nil
	}

	// 获取命令的第一个单词（主命令）
	mainCmd := parts[0]

	commandNames, err := t.repo.GetAllCommandNames()
	if err != nil {
		return "", nil, err
	}

	// 先检查主命令
	mainMatches := utils.MatchMultiple(mainCmd, commandNames, 5)

	// 如果主命令匹配度为100%，检查子命令
	if len(mainMatches) > 0 && mainMatches[0].Score == 1.0 {
		// 主命令正确，检查是否有子命令
		if len(parts) > 1 {
			subCmd := parts[1]

			// 获取该主命令的所有子命令
			subCommandNames, err := t.repo.GetAllCommandOptionNames(mainCmd)
			if err != nil {
				// 如果没有子命令，返回原始命令
				return command, nil, nil
			}

			if len(subCommandNames) > 0 {
				// 检查子命令的拼写
				subMatches := utils.MatchMultiple(subCmd, subCommandNames, 5)

				// 检查是否完全匹配
				isExactMatch := false
				for _, match := range subMatches {
					if match.Command == subCmd && match.Score == 1.0 {
						isExactMatch = true
						break
					}
				}

				// 如果子命令完全匹配，不返回任何建议
				if isExactMatch {
					return command, nil, nil
				}

				// 子命令不匹配或拼写错误，返回修正建议
				var results []types.MatchResult
				for _, match := range subMatches {
					// 构建新的命令：主命令 + 修正的子命令 + 剩余参数
					newCommand := mainCmd + " " + match.Command
					if len(parts) > 2 {
						newCommand += " " + strings.Join(parts[2:], " ")
					}

					// 获取子命令描述
					subCommands, _ := t.repo.GetCommandOptions(mainCmd)
					desc := ""
					for _, sc := range subCommands {
						if sc.OptionName == match.Command {
							desc = sc.Description
							break
						}
					}

					results = append(results, types.MatchResult{
						Command: newCommand,
						Score:   match.Score,
						Desc:    desc,
					})
				}

				if len(results) > 0 {
					return command, results, nil
				}
			}
		}
		return command, nil, nil
	}

	// 主命令可能有拼写错误，返回主命令的修正建议
	if len(mainMatches) > 0 && mainMatches[0].Score < 0.5 {
		_, err := t.repo.FindCommandByName(mainCmd)
		if err == nil {
			return command, nil, nil
		}
	}

	// 为主命令匹配结果添加描述
	for i := range mainMatches {
		cmd, err := t.repo.FindCommandByName(mainMatches[i].Command)
		if err != nil {
			mainMatches[i].Desc = ""
			continue
		}
		mainMatches[i].Desc = cmd.Description
	}

	return command, mainMatches, nil
}

// LLM impl

type LlmTypo struct {
	hs  scanner.IScanner
	llm openai.Client
}

var matchResultSchema = GenerateSchema[types.MatchResult]()

func NewLlmTypo(hs scanner.IScanner) ITypo {

	client := openai.NewClient(
		option.WithAPIKey("sk-1111"),
		option.WithBaseURL("https://api.openai.com/v1"),
	)

	return &LlmTypo{
		hs,
		client,
	}
}

func (l LlmTypo) Typo() (string, []types.MatchResult, error) {

	command, err := l.hs.Scan()
	if err != nil {
		return "", nil, err
	}

	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        "command_match_schema",
		Description: openai.String("The typo command match schema"),
		Schema:      matchResultSchema,
		Strict:      openai.Bool(true),
	}

	chatCompletion, err := l.llm.Chat.Completions.New(
		context.Background(),
		openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(systemPrompt()),
				openai.UserMessage(command),
			},
			ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
				OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{JSONSchema: schemaParam},
			},
			Model: "Qwen3-flash",
		},
	)
	if err != nil {
		return "", nil, err
	}

	return returnCommand(command, chatCompletion)
}

func returnCommand(cmd string, res *openai.ChatCompletion) (string, []types.MatchResult, error) {

	var results []types.MatchResult
	err := json.Unmarshal([]byte(res.Choices[0].Message.Content), &results)
	if err != nil {
		return "", nil, err
	}

	return cmd, results, nil
}

func systemPrompt() string {

	return `
Your role is to create a command-line tool that checks and fixes typos.
When a user enters a command, you need to check its correctness and provide suggestions for fixing it if necessary.

Note: During the fixing process, think carefully:

- Do not arbitrarily change the name or meaning of the command;

- Do not add new options to the input command string;

- Do not delete parameter information from the original command string;

- Always remember, it's just a typo; don't overcomplicate things.

You need to return a JSON-formatted array string containing five possible command fix candidates, with the following fields:

- command: The corrected command. Note that this may just be a typo; do not arbitrarily change the name or meaning of the command;

- score: The confidence range of the corrected command, represented as a floating-point number between 0 and 1, providing suggestions to the user;

- desc: A description of the command, indicating what the current command can do.

Below is an example of input and output using ‘gti st’:

input：

gti st

output：

[
  {
    "command": "git status",
    "score": 0.9,
    "desc": "查看当前仓库状态"
  },
  {
    "command": "git stash",
    "score": 0.8,
    "desc": "暂存当前工作区状态"
  },
  {
    "command": "git stash list",
    "score": 0.7,
    "desc": "查看暂存列表"
  },
  {
    "command": "git stash apply",
    "score": 0.6,
    "desc": "应用暂存的工作区状态"
  },
  {
    "command": "git stash drop",
    "score": 0.5,
    "desc": "删除暂存的工作区状态"
  }
]

Below is an example of the input and output of the ‘git branhc’ (branch misspelling) command:

input：

git branhc

output：

[
  {
    "command": "git branch",
    "score": 0.9,
    "desc": "查看本地分支"
  },
  {
    "command": "git checkout",
    "score": 0.8,
    "desc": "切换到指定分支"
  },
  {
    "command": "git branch -d",
    "score": 0.7,
    "desc": "删除本地分支"
  },
  {
    "command": "git branch -D",
    "score": 0.6,
    "desc": "强制删除本地分支"
  },
  {
    "command": "git branch -m",
    "score": 0.5,
    "desc": "重命名本地分支"
  }
]
  `
}

func GenerateSchema[T any]() interface{} {

	// Structured Outputs uses a subset of JSON schema
	// These flags are necessary to comply with the subset
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}

	var v T
	schema := reflector.Reflect(v)
	return schema
}
