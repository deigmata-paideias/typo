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
	// Typo returns the corrected command candidate list and original command
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

// Typo returns the corrected command candidate list and original command
// first return: original command
// second return: candidate commands with scores
func (t *LocalTypo) Typo() (string, []types.MatchResult, error) {

	command, err := t.hs.Scan()
	if err != nil {
		return "", nil, err
	}

	// Split command
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return command, nil, nil
	}

	// Get the first word of the command (main command)
	mainCmd := parts[0]

	commandNames, err := t.repo.GetAllCommandNames()
	if err != nil {
		return "", nil, err
	}

	// Check main command first
	mainMatches := utils.MatchMultiple(mainCmd, commandNames, 5)

	var allResults []types.MatchResult

	// If main command matches 100%, check subcommands
	if len(mainMatches) > 0 && mainMatches[0].Score == 1.0 {
		// Main command is correct, check if there are subcommands
		if len(parts) > 1 {
			subCmd := parts[1]

			// Get all subcommands for this main command
			subCommandNames, err := t.repo.GetAllCommandOptionNames(mainCmd)
			if err != nil {
				// If there are no subcommands, return the original command
				return command, nil, nil
			}

			if len(subCommandNames) > 0 {
				// Check subcommand spelling
				subMatches := utils.MatchMultiple(subCmd, subCommandNames, 5)

				// Check if there's an exact match
				isExactMatch := false
				for _, match := range subMatches {
					if match.Command == subCmd && match.Score == 1.0 {
						isExactMatch = true
						break
					}
				}

				// If subcommand matches exactly, don't return any suggestions
				if isExactMatch {
					return command, nil, nil
				}

				// Subcommand doesn't match or has spelling errors, return correction suggestions
				var results []types.MatchResult
				for _, match := range subMatches {
					// Build new command: main command + corrected subcommand + remaining arguments
					newCommand := mainCmd + " " + match.Command
					if len(parts) > 2 {
						newCommand += " " + strings.Join(parts[2:], " ")
					}

					// Get subcommand description
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

	// Main command has spelling errors
	// Strategy: Try to match both main command and subcommand (if exists) combinations
	if len(parts) > 1 {
		subCmd := parts[1]

		// For each possible main command match, try to find subcommand matches
		for _, mainMatch := range mainMatches {
			if mainMatch.Score < 0.5 {
				continue // Skip very low similarity matches
			}

			// Get subcommands for the matched main command
			subCommandNames, err := t.repo.GetAllCommandOptionNames(mainMatch.Command)
			if err == nil && len(subCommandNames) > 0 {
				// Try to match the subcommand
				subMatches := utils.MatchMultiple(subCmd, subCommandNames, 3)

				for _, subMatch := range subMatches {
					if subMatch.Score < 0.5 {
						continue
					}

					// Calculate combined score (weighted average)
					combinedScore := (mainMatch.Score * 0.6) + (subMatch.Score * 0.4)

					// Build corrected command
					newCommand := mainMatch.Command + " " + subMatch.Command
					if len(parts) > 2 {
						newCommand += " " + strings.Join(parts[2:], " ")
					}

					// Get description
					subCommands, _ := t.repo.GetCommandOptions(mainMatch.Command)
					desc := ""
					for _, sc := range subCommands {
						if sc.OptionName == subMatch.Command {
							desc = sc.Description
							break
						}
					}

					allResults = append(allResults, types.MatchResult{
						Command: newCommand,
						Score:   combinedScore,
						Desc:    desc,
					})
				}
			}
		}
	}

	// Main command might have spelling errors, return main command correction suggestions
	if len(mainMatches) > 0 && mainMatches[0].Score < 0.5 {
		_, err := t.repo.FindCommandByName(mainCmd)
		if err == nil {
			return command, nil, nil
		}
	}

	// Add descriptions for main command match results
	for i := range mainMatches {
		cmd, err := t.repo.FindCommandByName(mainMatches[i].Command)
		if err != nil {
			mainMatches[i].Desc = ""
			continue
		}
		mainMatches[i].Desc = cmd.Description

		// If there are parts after main command, append them to the suggestion
		if len(parts) > 1 {
			mainMatches[i].Command = mainMatches[i].Command + " " + strings.Join(parts[1:], " ")
		}
	}

	// Combine results: prioritize combined matches, then main command matches
	if len(allResults) > 0 {
		allResults = append(allResults, mainMatches...)
		// Sort by score and limit results
		return command, utils.SortAndLimitResults(allResults, 5), nil
	}

	return command, mainMatches, nil
}

// LLM impl

type LlmTypo struct {
	hs    scanner.IScanner
	llm   openai.Client
	model string
}

var matchResultSchema = GenerateSchema[[]types.MatchResult]()

func NewLlmTypo(hs scanner.IScanner, apiKey, baseURL, model string) ITypo {

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
		option.WithBaseURL(baseURL),
	)

	return &LlmTypo{
		hs:    hs,
		llm:   client,
		model: model,
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
			Model: l.model,
		},
	)
	if err != nil {
		return "", nil, err
	}

	return returnCommand(command, chatCompletion)
}

func returnCommand(cmd string, res *openai.ChatCompletion) (string, []types.MatchResult, error) {

	var results []types.MatchResult
	// debug: println(res.Choices[0].Message.Content)
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
    "desc": "Check current repository status"
  },
  {
    "command": "git stash",
    "score": 0.8,
    "desc": "Stash current working directory state"
  },
  {
    "command": "git stash list",
    "score": 0.7,
    "desc": "View stash list"
  },
  {
    "command": "git stash apply",
    "score": 0.6,
    "desc": "Apply stashed working directory state"
  },
  {
    "command": "git stash drop",
    "score": 0.5,
    "desc": "Delete stashed working directory state"
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
    "desc": "View local branches"
  },
  {
    "command": "git checkout",
    "score": 0.8,
    "desc": "Switch to specified branch"
  },
  {
    "command": "git branch -d",
    "score": 0.7,
    "desc": "Delete local branch"
  },
  {
    "command": "git branch -D",
    "score": 0.6,
    "desc": "Force delete local branch"
  },
  {
    "command": "git branch -m",
    "score": 0.5,
    "desc": "Rename local branch"
  }
]
  `
}

func GenerateSchema[T any]() any {

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

// NewTypo creates a Typo instance based on configuration
func NewTypo(config *types.Config, hs scanner.IScanner, repo repository.IRepository) ITypo {

	switch config.Mode {
	case types.LLM:
		return NewLlmTypo(hs, config.LLM.ApiKey, config.LLM.BaseUrl, config.LLM.Model)
	case types.Local:
		fallthrough
	default:
		return NewLocalTypo(repo, hs)
	}
}
