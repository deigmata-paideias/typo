package git

import (
	"strings"
)

type GitCommitResetRule struct{}

func (r *GitCommitResetRule) ID() string {
	return "git_commit_reset"
}

func (r *GitCommitResetRule) Match(command string, output string) bool {
    // Matches any git commit command
    return strings.Contains(command, "git commit")
}

func (r *GitCommitResetRule) GetNewCommand(command string, output string) string {
    return "git reset HEAD~"
}
