package git

import (
	"strings"
)

type GitStashRule struct{}

func (r *GitStashRule) ID() string { return "git_stash" }

func (r *GitStashRule) Match(command string, output string) bool {
	return strings.Contains(output, "or stash them")
}

func (r *GitStashRule) GetNewCommand(command string, output string) string {
	return "git stash && " + command
}
