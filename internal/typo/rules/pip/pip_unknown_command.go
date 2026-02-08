package pip

import (
	"regexp"
	"strings"
)

type PipUnknownCommandRule struct{}

func (r *PipUnknownCommandRule) ID() string { return "pip_unknown_command" }

func (r *PipUnknownCommandRule) Match(command string, output string) bool {
	return strings.Contains(command, "pip") &&
		strings.Contains(output, "unknown command") &&
		strings.Contains(output, "maybe you meant")
}

func (r *PipUnknownCommandRule) GetNewCommand(command string, output string) string {
	reBroken := regexp.MustCompile(`ERROR: unknown command "([^"]+)"`)
	reSuggestion := regexp.MustCompile(`maybe you meant "([^"]+)"`)

	brokenMatches := reBroken.FindStringSubmatch(output)
	suggestionMatches := reSuggestion.FindStringSubmatch(output)

	if len(brokenMatches) > 1 && len(suggestionMatches) > 1 {
		return strings.Replace(command, brokenMatches[1], suggestionMatches[1], 1)
	}
	return command
}
