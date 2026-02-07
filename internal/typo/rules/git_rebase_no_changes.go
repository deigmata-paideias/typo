package rules

import (
	"strings"
)

type GitRebaseNoChangesRule struct{}

func (r *GitRebaseNoChangesRule) ID() string {
	return "git_rebase_no_changes"
}

func (r *GitRebaseNoChangesRule) Match(command string, output string) bool {
    return strings.Contains(command, "rebase") && strings.Contains(command, "--continue") &&
           strings.Contains(output, "No changes - did you forget to use 'git add'?")
}

func (r *GitRebaseNoChangesRule) GetNewCommand(command string, output string) string {
    return "git rebase --skip"
}
