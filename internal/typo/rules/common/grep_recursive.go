package common

import (
	"strings"
)

type GrepRecursiveRule struct{}

func (r *GrepRecursiveRule) ID() string { return "grep_recursive" }

func (r *GrepRecursiveRule) Match(command string, output string) bool {
	return strings.Contains(command, "grep") && strings.Contains(strings.ToLower(output), "is a directory")
}

func (r *GrepRecursiveRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "grep", "grep -r", 1)
}
