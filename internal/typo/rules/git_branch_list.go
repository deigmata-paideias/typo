package rules

import (
	"strings"
)

type GitBranchListRule struct{}

func (r *GitBranchListRule) ID() string {
	return "git_branch_list"
}

func (r *GitBranchListRule) Match(command string, output string) bool {
	parts := strings.Fields(command)
	// git branch list
	if len(parts) >= 3 && parts[1] == "branch" && parts[2] == "list" {
		return true
	}
	return false
}

func (r *GitBranchListRule) GetNewCommand(command string, output string) string {
	return "git branch --delete list && git branch"
}
