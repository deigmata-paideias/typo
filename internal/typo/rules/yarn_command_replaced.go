package rules

import (
	"regexp"
	"strings"
)

type YarnCommandReplacedRule struct{}

func (r *YarnCommandReplacedRule) ID() string { return "yarn_command_replaced" }

func (r *YarnCommandReplacedRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "yarn") &&
		strings.Contains(output, "Run \"") &&
		strings.Contains(output, "\" instead")
}

func (r *YarnCommandReplacedRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`Run "(.*)" instead`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		return matches[1]
	}
	return command
}
