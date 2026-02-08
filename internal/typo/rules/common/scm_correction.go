package common

import (
	"os"
	"strings"
)

type ScmCorrectionRule struct{}

func (r *ScmCorrectionRule) ID() string { return "scm_correction" }

func (r *ScmCorrectionRule) Match(command string, output string) bool {
	// git: fatal: Not a git repository
	// hg: abort: no repository found
	if strings.HasPrefix(command, "git") && strings.Contains(output, "fatal: Not a git repository") {
		// check if .hg exists
		if _, err := os.Stat(".hg"); err == nil {
			return true
		}
	}
	if strings.HasPrefix(command, "hg") && strings.Contains(output, "abort: no repository found") {
		// check if .git exists
		if _, err := os.Stat(".git"); err == nil {
			return true
		}
	}
	return false
}

func (r *ScmCorrectionRule) GetNewCommand(command string, output string) string {
	parts := strings.Fields(command)
	if strings.HasPrefix(command, "git") {
		parts[0] = "hg"
	} else {
		parts[0] = "git"
	}
	return strings.Join(parts, " ")
}
