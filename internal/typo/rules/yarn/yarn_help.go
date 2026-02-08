package yarn

import (
	"regexp"
	"strings"
)

type YarnHelpRule struct{}

func (r *YarnHelpRule) ID() string { return "yarn_help" }

func (r *YarnHelpRule) Match(command string, output string) bool {
	// yarn help <command> -> Visit <url> for documentation
	return strings.HasPrefix(command, "yarn") &&
		strings.Contains(command, "help") &&
		strings.Contains(output, "Visit ") &&
		strings.Contains(output, "for documentation about this command")
}

func (r *YarnHelpRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`Visit ([^ ]*) for documentation about this command.`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		url := matches[1]
		// macOS
		return "open " + url
	}
	return command
}
