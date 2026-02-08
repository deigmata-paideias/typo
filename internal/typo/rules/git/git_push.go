package git

import (
	"regexp"
	"strings"
)

type GitPushRule struct{}

func (r *GitPushRule) ID() string {
	return "git_push"
}

func (r *GitPushRule) Match(command string, output string) bool {
	return strings.Contains(command, "push") &&
		strings.Contains(output, "git push --set-upstream")
}

func (r *GitPushRule) GetNewCommand(command string, output string) string {
	// Re-implementation of logic

	re := regexp.MustCompile(`git push (.*)`)
	matches := re.FindAllStringSubmatch(output, -1)
	if len(matches) > 0 {
		lastMatch := matches[len(matches)-1]
		if len(lastMatch) > 1 {
			args := strings.TrimSpace(lastMatch[1])
			return "git push " + args
		}
	}

	return command
}
