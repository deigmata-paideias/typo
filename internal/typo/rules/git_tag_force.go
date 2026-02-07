package rules

import (
	"strings"
)

type GitTagForceRule struct{}

func (r *GitTagForceRule) ID() string { return "git_tag_force" }

func (r *GitTagForceRule) Match(command string, output string) bool {
	return strings.Contains(command, "tag") && strings.Contains(output, "already exists")
}

func (r *GitTagForceRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "tag", "tag --force", 1)
}
