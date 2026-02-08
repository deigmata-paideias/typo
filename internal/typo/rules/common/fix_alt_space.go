package common

import (
	"strings"
)

type FixAltSpaceRule struct{}

func (r *FixAltSpaceRule) ID() string { return "fix_alt_space" }

func (r *FixAltSpaceRule) Match(command string, output string) bool {
	// Non-breaking space is \xc2\xa0 or \u00a0
	if strings.Contains(output, "command not found") {
		if strings.Contains(command, "\u00a0") || strings.Contains(command, "\xc2\xa0") {
			return true
		}
	}
	return false
}

func (r *FixAltSpaceRule) GetNewCommand(command string, output string) string {
	s := strings.ReplaceAll(command, "\u00a0", " ")
	s = strings.ReplaceAll(s, "\xc2\xa0", " ")
	return s
}
