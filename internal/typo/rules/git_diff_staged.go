package rules

import (
	"strings"
)

type GitDiffStagedRule struct{}

func (r *GitDiffStagedRule) ID() string {
	return "git_diff_staged"
}

func (r *GitDiffStagedRule) Match(command string, output string) bool {
	// git diff ...
	if !strings.Contains(command, "git diff") {
		return false
	}
	// Only if not already staged
	if strings.Contains(command, "--staged") || strings.Contains(command, "--cached") {
		return false
	}
	return true
}

func (r *GitDiffStagedRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "diff", "diff --staged", 1)
}
