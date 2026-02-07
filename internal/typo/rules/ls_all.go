package rules

import (
	"strings"
)

type LsAllRule struct{}

func (r *LsAllRule) ID() string {
	return "ls_all"
}

func (r *LsAllRule) Match(command string, output string) bool {
	// Match if command is likely 'ls' and output is empty
	// Note: output might contain whitespace
	return strings.HasPrefix(command, "ls") && strings.TrimSpace(output) == ""
}

func (r *LsAllRule) GetNewCommand(command string, output string) string {
	// ls -> ls -A
	// ls dir -> ls -A dir

	// Just insert -A after ls
	// Naive replacement of first occurrences
	return strings.Replace(command, "ls", "ls -A", 1)
}
