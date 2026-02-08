package common

import (
	"strings"
)

type LsLahRule struct{}

func (r *LsLahRule) ID() string { return "ls_lah" }

func (r *LsLahRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "ls") && !strings.Contains(command, "ls -")
}

func (r *LsLahRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "ls", "ls -lah", 1)
}
