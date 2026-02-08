package git

import (
	"strings"
)

type GitRebaseMergeDirRule struct{}

func (r *GitRebaseMergeDirRule) ID() string {
	return "git_rebase_merge_dir"
}

func (r *GitRebaseMergeDirRule) Match(command string, output string) bool {
    return strings.Contains(command, "rebase") &&
           strings.Contains(output, "It seems that there is already a rebase-merge directory") &&
           strings.Contains(output, "I wonder if you are in the middle of another rebase")
}

func (r *GitRebaseMergeDirRule) GetNewCommand(command string, output string) string {
    // Only returning one option for now.
    return "git rebase --continue"
}
