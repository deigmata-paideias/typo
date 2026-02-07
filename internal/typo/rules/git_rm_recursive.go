package rules

import (
	"strings"
)

type GitRmRecursiveRule struct{}

func (r *GitRmRecursiveRule) ID() string { return "git_rm_recursive" }

func (r *GitRmRecursiveRule) Match(command string, output string) bool {
	return strings.Contains(command, "git rm") &&
		strings.Contains(output, "fatal: not removing '") &&
		strings.Contains(output, "' recursively without -r")
}

func (r *GitRmRecursiveRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "git rm", "git rm -r", 1)
}
