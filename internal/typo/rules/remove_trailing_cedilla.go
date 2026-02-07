package rules

import (
	"strings"
)

type RemoveTrailingCedillaRule struct{}

func (r *RemoveTrailingCedillaRule) ID() string { return "remove_trailing_cedilla" }

const cedilla = "ç"

func (r *RemoveTrailingCedillaRule) Match(command string, output string) bool {
	return strings.HasSuffix(command, cedilla)
}

func (r *RemoveTrailingCedillaRule) GetNewCommand(command string, output string) string {
	return strings.TrimSuffix(command, cedilla)
}
