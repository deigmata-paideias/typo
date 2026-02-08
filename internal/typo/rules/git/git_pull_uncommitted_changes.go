package git

import (
	"strings"
)

type GitPullUncommittedChangesRule struct{}

func (r *GitPullUncommittedChangesRule) ID() string {
	return "git_pull_uncommitted_changes"
}

func (r *GitPullUncommittedChangesRule) Match(command string, output string) bool {
	return strings.Contains(command, "pull") &&
		(strings.Contains(output, "You have unstaged changes") ||
			strings.Contains(output, "contains uncommitted changes") ||
			strings.Contains(output, "Please commit your changes or stash them before you"))
}

func (r *GitPullUncommittedChangesRule) GetNewCommand(command string, output string) string {
	return "git stash && " + command + " && git stash pop"
}
