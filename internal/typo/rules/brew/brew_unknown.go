package brew

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type BrewUnknownCommandRule struct{}

func (r *BrewUnknownCommandRule) ID() string {
	return "brew_unknown_command"
}

func (r *BrewUnknownCommandRule) Match(command string, output string) bool {
	return strings.Contains(command, "brew") &&
		strings.Contains(output, "Unknown command")
}

func (r *BrewUnknownCommandRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`Unknown command: ([a-z]+)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return command
	}
	brokenCmd := matches[1]

	brewCmds := []string{
		"install", "uninstall", "search", "list", "update", "upgrade",
		"info", "home", "doctor", "edit", "cleanup", "tap", "services",
	}

	bestMatch := utils.Match(brokenCmd, brewCmds)
	if bestMatch != "" {
		return strings.Replace(command, brokenCmd, bestMatch, 1)
	}

	return command
}
