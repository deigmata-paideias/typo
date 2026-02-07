package rules

import (
	"strings"
)

type GitCommitAddRule struct{}

func (r *GitCommitAddRule) ID() string {
	return "git_commit_add"
}

func (r *GitCommitAddRule) Match(command string, output string) bool {
	return strings.Contains(command, "commit") &&
		strings.Contains(output, "no changes added to commit")
}

func (r *GitCommitAddRule) GetNewCommand(command string, output string) string {
	// suggest -a
	// git commit -a ... (command args?)
	// command is "git commit -m 'foo'" -> "git commit -a -m 'foo'"

	// replace "commit" with "commit -a"
	return strings.Replace(command, "commit", "commit -a", 1)
}
