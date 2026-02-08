package common

import (
	"strings"
)

type DryRule struct{}

func (r *DryRule) ID() string {
	return "dry"
}

func (r *DryRule) Match(command string, output string) bool {
	parts := strings.Fields(command)
	if len(parts) >= 2 {
		return parts[0] == parts[1]
	}
	return false
}

func (r *DryRule) GetNewCommand(command string, output string) string {
	// git git push -> git push
	parts := strings.Fields(command)
	return strings.Join(parts[1:], " ")
}
