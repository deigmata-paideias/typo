package git

import (
	"strings"
)

type GitStashPopRule struct{}

func (r *GitStashPopRule) ID() string { return "git_stash_pop" }

func (r *GitStashPopRule) Match(command string, output string) bool {
	return strings.Contains(command, "stash") && strings.Contains(command, "pop") &&
		strings.Contains(output, "Your local changes to the following files would be overwritten by merge")
}

func (r *GitStashPopRule) GetNewCommand(command string, output string) string {
	return "git add --update && git stash pop && git reset ."
}
