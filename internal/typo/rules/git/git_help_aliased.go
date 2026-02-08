package git

import (
	"regexp"
	"strings"
)

type GitHelpAliasedRule struct{}

func (r *GitHelpAliasedRule) ID() string {
	return "git_help_aliased"
}

func (r *GitHelpAliasedRule) Match(command string, output string) bool {
	return strings.Contains(command, "help") && strings.Contains(output, " is aliased to ")
}

func (r *GitHelpAliasedRule) GetNewCommand(command string, output string) string {
	// `checkout` is aliased to `...`
	// Extract aliased command
	// Python code parses carefully. output example: "`st` is aliased to `status`" or similar?
	// "git help st" -> "`st` is aliased to `status`"

	// Regex: `([^`]*)` is aliased to `([^`]*)`  <-- checking format

	re := regexp.MustCompile("`([^`]*)` is aliased to `([^`]*)`")
	matches := re.FindStringSubmatch(output)
	if len(matches) > 2 {
		realCmd := matches[2]
		// Extract first word of alias? "status -s" -> "status"
		realCmd = strings.Split(realCmd, " ")[0]

		return "git help " + realCmd
	}

	// Python split fallback
	// aliased = command.output.split('`', 2)[2].split("'", 1)[0].split(' ', 1)[0]
	return command
}
