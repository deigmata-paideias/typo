package typo

import (
	"context"
	"encoding/json"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/invopop/jsonschema"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"

	"github.com/deigmata-paideias/typo/internal/repository"
	"github.com/deigmata-paideias/typo/internal/scanner"
	"github.com/deigmata-paideias/typo/internal/types"
	"github.com/deigmata-paideias/typo/internal/typo/rules"
	"github.com/deigmata-paideias/typo/internal/typo/rules/adb"
	"github.com/deigmata-paideias/typo/internal/typo/rules/ag"
	"github.com/deigmata-paideias/typo/internal/typo/rules/apt"
	"github.com/deigmata-paideias/typo/internal/typo/rules/aws"
	"github.com/deigmata-paideias/typo/internal/typo/rules/az"
	"github.com/deigmata-paideias/typo/internal/typo/rules/brew"
	"github.com/deigmata-paideias/typo/internal/typo/rules/cargo"
	"github.com/deigmata-paideias/typo/internal/typo/rules/common"
	"github.com/deigmata-paideias/typo/internal/typo/rules/composer"
	"github.com/deigmata-paideias/typo/internal/typo/rules/conda"
	"github.com/deigmata-paideias/typo/internal/typo/rules/django"
	"github.com/deigmata-paideias/typo/internal/typo/rules/dnf"
	"github.com/deigmata-paideias/typo/internal/typo/rules/docker"
	"github.com/deigmata-paideias/typo/internal/typo/rules/fab"
	"github.com/deigmata-paideias/typo/internal/typo/rules/git"
	"github.com/deigmata-paideias/typo/internal/typo/rules/golang"
	"github.com/deigmata-paideias/typo/internal/typo/rules/gradle"
	"github.com/deigmata-paideias/typo/internal/typo/rules/grunt"
	"github.com/deigmata-paideias/typo/internal/typo/rules/heroku"
	"github.com/deigmata-paideias/typo/internal/typo/rules/hosts"
	"github.com/deigmata-paideias/typo/internal/typo/rules/ifconfig"
	"github.com/deigmata-paideias/typo/internal/typo/rules/java"
	"github.com/deigmata-paideias/typo/internal/typo/rules/kubectl"
	"github.com/deigmata-paideias/typo/internal/typo/rules/mercurial"
	"github.com/deigmata-paideias/typo/internal/typo/rules/mvn"
	"github.com/deigmata-paideias/typo/internal/typo/rules/nixos"
	"github.com/deigmata-paideias/typo/internal/typo/rules/npm"
	"github.com/deigmata-paideias/typo/internal/typo/rules/pacman"
	"github.com/deigmata-paideias/typo/internal/typo/rules/php"
	"github.com/deigmata-paideias/typo/internal/typo/rules/pip"
	"github.com/deigmata-paideias/typo/internal/typo/rules/python"
	"github.com/deigmata-paideias/typo/internal/typo/rules/rails"
	"github.com/deigmata-paideias/typo/internal/typo/rules/react_native"
	"github.com/deigmata-paideias/typo/internal/typo/rules/systemctl"
	"github.com/deigmata-paideias/typo/internal/typo/rules/terraform"
	"github.com/deigmata-paideias/typo/internal/typo/rules/tmux"
	"github.com/deigmata-paideias/typo/internal/typo/rules/tsuru"
	"github.com/deigmata-paideias/typo/internal/typo/rules/vagrant"
	"github.com/deigmata-paideias/typo/internal/typo/rules/workon"
	"github.com/deigmata-paideias/typo/internal/typo/rules/yarn"
	"github.com/deigmata-paideias/typo/internal/typo/rules/yum"
	"github.com/deigmata-paideias/typo/internal/utils"
)

type ITypo interface {
	// Typo returns the corrected command candidate list and original command
	Typo() (string, []types.MatchResult, error)
}

// deduplicateResults removes duplicate commands, keeping the one with highest score
func deduplicateResults(results []types.MatchResult) []types.MatchResult {
	seen := make(map[string]types.MatchResult)

	for _, result := range results {
		if existing, found := seen[result.Command]; found {
			// Keep the one with higher score
			if result.Score > existing.Score {
				seen[result.Command] = result
			}
		} else {
			seen[result.Command] = result
		}
	}

	// Convert map back to slice
	deduplicated := make([]types.MatchResult, 0, len(seen))
	for _, result := range seen {
		deduplicated = append(deduplicated, result)
	}

	// Sort by score in descending order
	sort.Slice(deduplicated, func(i, j int) bool {
		return deduplicated[i].Score > deduplicated[j].Score
	})

	return deduplicated
}

type LocalTypo struct {
	repo  repository.IRepository
	hs    scanner.IScanner
	rules []rules.Rule
}

func NewLocalTypo(repo repository.IRepository, hs scanner.IScanner) ITypo {
	// Initialize standard rules
	r := []rules.Rule{
		&git.GitPushRule{},
		&git.GitCheckoutRule{},
		&git.GitBranchExistsRule{},
		&git.GitAddRule{},
		&common.MkdirPRule{},
		&common.CpOmittingDirectoryRule{},
		&common.RmDirRule{},
		&common.CdMkdirRule{},
		&common.CdParentRule{},
		&common.ChmodXRule{},
		&brew.BrewUnknownCommandRule{},
		&docker.DockerUnknownCommandRule{},
		&common.GrepArgumentsOrderRule{},
		&common.SudoRule{},
		&common.SedUnterminatedSRule{},
		&common.LsAllRule{},
		&common.CatDirRule{},
		&common.SlLsRule{},
		&common.TouchRule{},
		&golang.GoRunRule{},
		&common.ManNoSpaceRule{},
		&common.PythonCommandRule{},
		&common.GitDiffNoIndexRule{},
		&golang.GoUnknownCommandRule{},
		&adb.ADBUnknownCommandRule{},
		&ag.AgLiteralRule{},
		&apt.AptGetSearchRule{},
		&apt.AptListUpgradableRule{},
		&apt.AptUpgradeRule{},
		&apt.AptInvalidOperationRule{},
		&python.PythonExecuteRule{},
		&aws.AwsCliRule{},
		&az.AzCliRule{},
		&brew.BrewCaskDependencyRule{},
		&brew.BrewInstallRule{},
		&brew.BrewLinkRule{},
		&brew.BrewReinstallRule{},
		&brew.BrewUninstallRule{},
		&brew.BrewUpdateFormulaRule{},
		&cargo.CargoRule{},
		&cargo.CargoNoCommandRule{},
		&common.CdCorrectionRule{},
		&common.CdCsRule{},
		&common.ChocoInstallRule{},
		&composer.ComposerNotCommandRule{},
		&conda.CondaMistypeRule{},
		&common.CpCreateDestinationRule{},
		&common.Cpp11Rule{},
		&common.DirtyUntarRule{},
		&common.DirtyUnzipRule{},
		&django.DjangoSouthGhostRule{},
		&django.DjangoSouthMergeRule{},
		&dnf.DnfNoSuchCommandRule{},
		&docker.DockerImageUsedRule{},
		&docker.DockerLoginRule{},
		&common.DryRule{},
		&fab.FabCommandNotFoundRule{},
		&common.FixAltSpaceRule{},
		&common.FixFileRule{},
		&common.GemUnknownCommandRule{},
		&git.GitAddForceRule{},
		&git.GitBisectUsageRule{},
		&common.GitBranch0FlagRule{},
		&git.GitBranchDeleteRule{},
		&git.GitBranchDeleteCheckedOutRule{},
		&git.GitBranchListRule{},
		&common.GitCloneGitCloneRule{},
		&git.GitCloneMissingRule{},
		&git.GitCommitAddRule{},
		&common.GitCommitAmendRule{},
		&git.GitCommitResetRule{},
		&git.GitDiffStagedRule{},
		&common.GitFixStashRule{},
		&git.GitFlagAfterFilenameRule{},
		&git.GitHelpAliasedRule{},
		&git.GitHookBypassRule{},
		&git.GitLfsMistypeRule{},
		&common.GitMainMasterRule{},
		&common.GitMergeRule{},
		&git.GitMergeUnrelatedRule{},
		&git.GitNotCommandRule{},
		&common.GitPullRule{},
		&git.GitPullCloneRule{},
		&git.GitPullUncommittedChangesRule{},
		&git.GitPushDifferentBranchNamesRule{},
		&git.GitPushForceRule{},
		&common.GitPushPullRule{},
		&common.GitPushWithoutCommitsRule{},
		&git.GitRebaseMergeDirRule{},
		&git.GitRebaseNoChangesRule{},
		&git.GitRemoteDeleteRule{},
		&git.GitRemoteSeturlAddRule{},
		&common.GitRmLocalModificationsRule{},
		&git.GitRmRecursiveRule{},
		&git.GitRmStagedRule{},
		&git.GitStashRule{},
		&git.GitStashPopRule{},
		&git.GitTagForceRule{},
		&git.GitTwoDashesRule{},
		&common.GrepRecursiveRule{},
		&gradle.GradleWrapperRule{},
		&common.HasExistsScriptRule{},
		&heroku.HerokuNotCommandRule{},
		&hosts.HostsCliRule{},
		&gradle.GradleNoTaskRule{},
		&grunt.GruntTaskNotFoundRule{},
		&common.GulpNotTaskRule{},
		&ifconfig.IfconfigDeviceNotFoundRule{},
		&java.JavaRule{},
		&common.JavacRule{},
		&kubectl.KubectlGetMissingResourceRule{},
		&kubectl.KubectlMisspelledFlagRule{},
		&kubectl.KubectlNoContextRule{},
		&kubectl.KubectlResourceTypoRule{},
		&kubectl.KubectlUnknownCommandRule{},
		&common.LeinNotTaskRule{},
		&common.LnNoHardLinkRule{},
		&common.LnSOrderRule{},
		&common.LongFormHelpRule{},
		&common.LsLahRule{},
		&mercurial.MercurialRule{},
		&common.MissingSpaceBeforeSubcommandRule{},
		&mvn.MvnNoCommandRule{},
		&mvn.MvnUnknownLifecyclePhaseRule{},
		&nixos.NixosCmdNotFoundRule{},
		&common.NoSuchFileRule{},
		&common.NpmMissingScriptRule{},
		&common.NpmRunScriptRule{},
		&npm.NpmWrongCommandRule{},
		&common.OmnienvNoSuchCommandRule{},
		&common.OpenRule{},
		&common.PacmanInvalidOptionRule{},
		&pacman.PacmanNotFoundRule{},
		&pacman.PacmanRule{},
		&php.PhpSRule{},
		&pip.PipInstallRule{},
		&pip.PipUnknownCommandRule{},
		&common.PortAlreadyInUseRule{},
		&common.ProveRecursivelyRule{},
		&python.PythonModuleErrorRule{},
		&common.QuotationMarksRule{},
		&rails.RailsMigrationsPendingRule{},
		&react_native.ReactNativeCommandUnrecognizedRule{},
		&common.RemoveShellPromptLiteralRule{},
		&common.RemoveTrailingCedillaRule{},
		&common.RmRootRule{},
		&common.ScmCorrectionRule{},
		&common.SshKnownHostsRule{},
		&common.SudoCommandFromUserPathRule{},
		&common.SwitchLangRule{},
		&systemctl.SystemctlRule{},
		&terraform.TerraformInitRule{},
		&terraform.TerraformNoCommandRule{},
		&common.TestPyRule{},
		&tmux.TmuxRule{},
		&tsuru.TsuruLoginRule{},
		&tsuru.TsuruNotCommandRule{},
		&common.UnknownCommandRule{},
		&common.UnsudoRule{},
		&vagrant.VagrantUpRule{},
		&common.WhoisRule{},
		&workon.WorkonDoesntExistsRule{},
		&common.WrongHyphenBeforeSubcommandRule{},
		&yarn.YarnAliasRule{},
		&yarn.YarnCommandNotFoundRule{},
		&yarn.YarnCommandReplacedRule{},
		&yarn.YarnHelpRule{},
		&yum.YumInvalidOperationRule{},
	}
	return &LocalTypo{
		repo,
		hs,
		r,
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

	// --- Rules Check (Thefuck style) ---
	var ruleResults []types.MatchResult

	// Check if command is safe to re-run to capture output
	// We mimic 'thefuck' behavior which relies on command output for many rules
	safePrefixes := []string{"git", "mkdir", "brew", "ls", "cd", "grep", "cp", "cat", "echo", "touch", "sed", "docker", "rm", "sl", "cd..", "go", "man", "python", "adb", "ag", "aws", "cargo", "composer", "conda", "choco", "cs", "tar", "unzip", "g++", "clang++", "mv", "dnf", "fab", "gem", "gradle", "grunt", "gulp", "hostscli", "heroku", "ifconfig", "java", "javac", "lein", "ln", "hg", "mvn", "npm", "nix-env", "open", "xdg-open", "gnome-open", "kde-open", "goenv", "nodenv", "pyenv", "rbenv", "react-native", "rm", "sudo", "ssh", "scp", "systemctl", "terraform", "tmux", "tsuru", "vagrant", "whois", "pacman", "php", "pip", "pip2", "pip3", "prove", "workon", "yarn", "yum", "apt", "apt-get", "az"}
	isSafe := false
	for _, prefix := range safePrefixes {
		if strings.HasPrefix(command, prefix) {
			isSafe = true
			break
		}
	}

	if isSafe {
		// Run command to get output
		// We use zsh -c to run the full command line to capture stderr/stdout
		output, _ := utils.ExecCommandWithOutput("zsh", "-c", command)

		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		sysMessages := []string{
			"Nailed it.",
			"I got you.",
			"Fixed it for you.",
			"Don't worry, happens to the best of us.",
			"Let's try this instead.",
		}

		for _, rule := range t.rules {
			if rule.Match(command, output) {
				newCmd := rule.GetNewCommand(command, output)
				if newCmd != command {
					ruleResults = append(ruleResults, types.MatchResult{
						Command: newCmd,
						Score:   1.0,
						Desc:    sysMessages[rng.Intn(len(sysMessages))],
					})
				}
			}
		}
	}

	if len(ruleResults) > 0 {
		return command, deduplicateResults(ruleResults), nil
	}
	// --- End Rules Check ---

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
					return command, deduplicateResults(results), nil
				}
			}
		}
		return command, nil, nil
	}

	// Main command has spelling errors
	// Strategy: Try to match both main command and subcommand (if exists) combinations
	hasSubcommand := len(parts) > 1

	if hasSubcommand {
		subCmd := parts[1]

		// For each possible main command match, try to find subcommand matches
		for _, mainMatch := range mainMatches {
			if mainMatch.Score < 0.3 {
				continue // Skip very low similarity matches
			}

			// Get subcommands for the matched main command
			subCommandNames, err := t.repo.GetAllCommandOptionNames(mainMatch.Command)
			if err == nil && len(subCommandNames) > 0 {
				// Try to match the subcommand
				subMatches := utils.MatchMultiple(subCmd, subCommandNames, 5)

				for _, subMatch := range subMatches {
					if subMatch.Score < 0.3 {
						continue
					}

					// Calculate combined score using multiplication
					// This ensures both parts need good matches
					combinedScore := mainMatch.Score * subMatch.Score

					// Boost score for high-quality matches on both parts
					if mainMatch.Score >= 0.7 && subMatch.Score >= 0.7 {
						combinedScore = combinedScore * 1.3
						if combinedScore > 1.0 {
							combinedScore = 1.0
						}
					}

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
					if desc == "" {
						desc = mainMatch.Desc
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

	// Add descriptions for main command match results and rebuild full command with arguments
	var results []types.MatchResult
	for _, match := range mainMatches {
		// Build new command: corrected main command + remaining arguments
		newCommand := match.Command
		if len(parts) > 1 {
			newCommand += " " + strings.Join(parts[1:], " ")
		}

		cmd, err := t.repo.FindCommandByName(match.Command)
		desc := ""
		if err == nil {
			desc = cmd.Description
		}

		results = append(results, types.MatchResult{
			Command: newCommand,
			Score:   match.Score,
			Desc:    desc,
		})
	}

	return command, deduplicateResults(results), nil
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
