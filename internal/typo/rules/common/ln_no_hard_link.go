package common

import (
	"strings"
)

type LnNoHardLinkRule struct{}

func (r *LnNoHardLinkRule) ID() string { return "ln_no_hard_link" }

func (r *LnNoHardLinkRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "ln ") &&
		strings.Contains(output, "hard link not allowed for directory")
}

func (r *LnNoHardLinkRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "ln ", "ln -s ", 1)
}
