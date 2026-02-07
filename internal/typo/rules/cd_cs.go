package rules

import (
	"strings"
)

type CdCsRule struct{}

func (r *CdCsRule) ID() string {
	return "cd_cs"
}

func (r *CdCsRule) Match(command string, output string) bool {
	// cs foo
	return strings.HasPrefix(command, "cs ") || command == "cs"
}

func (r *CdCsRule) GetNewCommand(command string, output string) string {
	// cs foo -> cd foo
	// cs -> cd
	if command == "cs" {
		return "cd"
	}
	return "cd" + command[2:]
}
