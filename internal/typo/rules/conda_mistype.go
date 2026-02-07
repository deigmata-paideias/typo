package rules

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type CondaMistypeRule struct{}

func (r *CondaMistypeRule) ID() string { return "conda_mistype" }

func (r *CondaMistypeRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "conda") &&
		strings.Contains(output, "CommandNotFoundError")
}

func (r *CondaMistypeRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`CommandNotFoundError: No command 'conda ([^']+)'.`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return command
	}
	typo := matches[1]

	validCommands := []string{
		"clean", "config", "create", "help", "info", "install", "list", "package",
		"remove", "uninstall", "search", "update", "upgrade",
	}

	best := utils.Match(typo, validCommands)
	if best != "" {
		return strings.Replace(command, typo, best, 1)
	}
	return command
}
