package rules

import (
	"strings"
)

type GitMergeUnrelatedRule struct{}

func (r *GitMergeUnrelatedRule) ID() string {
	return "git_merge_unrelated"
}

func (r *GitMergeUnrelatedRule) Match(command string, output string) bool {
    return strings.Contains(command, "merge") &&
           strings.Contains(output, "fatal: refusing to merge unrelated histories")
}

func (r *GitMergeUnrelatedRule) GetNewCommand(command string, output string) string {
    return command + " --allow-unrelated-histories"
}
