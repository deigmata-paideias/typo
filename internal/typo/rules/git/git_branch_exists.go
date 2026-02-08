package git

import (
	"regexp"
	"strings"
)

type GitBranchExistsRule struct{}

func (r *GitBranchExistsRule) ID() string {
	return "git_branch_exists"
}

func (r *GitBranchExistsRule) Match(command string, output string) bool {
	// git branch foo
	// fatal: A branch named 'foo' already exists.
	return strings.Contains(output, "fatal: A branch named") &&
		strings.Contains(output, "already exists")
}

func (r *GitBranchExistsRule) GetNewCommand(command string, output string) string {
	// Extract branch name
	re := regexp.MustCompile(`fatal: A branch named '(.+)' already exists.`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return command
	}
	branchName := matches[1]

	// If I tried to create a branch 'foo' and it exists, maybe I meant to checkout it?
	// git branch foo -> git checkout foo
	// git branch -c foo -> git checkout -b foo?? (If it exists, git checkout foo)

	// Simple fix: suggest checkout
	return "git checkout " + branchName
}
