package rules

import (
	"regexp"
	"strings"
)

type YarnAliasRule struct{}

func (r *YarnAliasRule) ID() string { return "yarn_alias" }

func (r *YarnAliasRule) Match(command string, output string) bool {
	// yarn <something> -> Did you mean `alias`?
	return strings.HasPrefix(command, "yarn") && strings.Contains(output, "Did you mean")
}

func (r *YarnAliasRule) GetNewCommand(command string, output string) string {
	// Did you mean `list`?
	re := regexp.MustCompile("Did you mean [`\"](?:yarn )?([^`\"]*)[`\"]")
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		fix := matches[1]
		// Replace the subcommand.
		// command: yarn lst
		parts := strings.Fields(command)
		if len(parts) > 1 {
			return strings.Replace(command, parts[1], fix, 1)
		}
	}
	return command
}
