package common

import (
	"strings"
)

type UnsudoRule struct{}

func (r *UnsudoRule) ID() string { return "unsudo" }

func (r *UnsudoRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "sudo") &&
		strings.Contains(strings.ToLower(output), "you cannot perform this operation as root")
}

func (r *UnsudoRule) GetNewCommand(command string, output string) string {
	return strings.TrimPrefix(command, "sudo ")
}
