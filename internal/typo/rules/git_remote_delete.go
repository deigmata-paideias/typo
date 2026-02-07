package rules

import (
	"strings"
)

type GitRemoteDeleteRule struct{}

func (r *GitRemoteDeleteRule) ID() string { return "git_remote_delete" }

func (r *GitRemoteDeleteRule) Match(command string, output string) bool {
	return strings.Contains(command, "git remote delete") &&
		strings.Contains(output, "error: unknown subcommand: `delete'")
}

func (r *GitRemoteDeleteRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "remote delete", "remote remove", 1)
}
