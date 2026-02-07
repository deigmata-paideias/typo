package rules

import (
	"strings"
)

type GitPullCloneRule struct{}

func (r *GitPullCloneRule) ID() string {
	return "git_pull_clone"
}

func (r *GitPullCloneRule) Match(command string, output string) bool {
    return strings.Contains(output, "fatal: Not a git repository") &&
           strings.Contains(output, "Stopping at filesystem boundary (GIT_DISCOVERY_ACROSS_FILESYSTEM not set).")
}

func (r *GitPullCloneRule) GetNewCommand(command string, output string) string {
    return strings.Replace(command, "pull", "clone", 1)
}
