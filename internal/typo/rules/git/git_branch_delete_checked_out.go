package git

import (
	"strings"
)

type GitBranchDeleteCheckedOutRule struct{}

func (r *GitBranchDeleteCheckedOutRule) ID() string {
	return "git_branch_delete_checked_out"
}

func (r *GitBranchDeleteCheckedOutRule) Match(command string, output string) bool {
	// Check branch -d or branch -D matches
	// Note: command string might have different spacing, but Contains should work.
	isDelete := strings.Contains(command, "branch -d") || strings.Contains(command, "branch -D")

	return isDelete &&
		strings.Contains(output, "error: Cannot delete branch '") &&
		strings.Contains(output, "' checked out at '")
}

func (r *GitBranchDeleteCheckedOutRule) GetNewCommand(command string, output string) string {
	// switch to master and force delete
	// replace -d with -D just in case
	cmd := strings.Replace(command, "-d", "-D", 1)
	return "git checkout master && " + cmd
}
