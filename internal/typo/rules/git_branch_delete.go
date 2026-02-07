package rules

import (
	"strings"
)

type GitBranchDeleteRule struct{}

func (r *GitBranchDeleteRule) ID() string {
	return "git_branch_delete"
}

func (r *GitBranchDeleteRule) Match(command string, output string) bool {
    return strings.Contains(command, "branch -d") &&
           strings.Contains(output, "If you are sure you want to delete it")
}

func (r *GitBranchDeleteRule) GetNewCommand(command string, output string) string {
    return strings.Replace(command, "-d", "-D", 1)
}
