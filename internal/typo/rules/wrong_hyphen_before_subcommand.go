package rules

import (
	"os/exec"
	"strings"
)

type WrongHyphenBeforeSubcommandRule struct{}

func (r *WrongHyphenBeforeSubcommandRule) ID() string { return "wrong_hyphen_before_subcommand" }

func (r *WrongHyphenBeforeSubcommandRule) Match(command string, output string) bool {
	// git-push -> git push
	// if git-push fails (or not executable) and git exists.
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return false
	}
	first := parts[0]
	if !strings.Contains(first, "-") {
		return false
	}

	// Check if first part is an executable
	if _, err := exec.LookPath(first); err == nil {
		return false // It is a valid executable
	}

	// Split by hypen
	// git-push -> git, push
	subParts := strings.SplitN(first, "-", 2)
	if len(subParts) < 2 {
		return false
	}
	bin := subParts[0]

	// Check if bin is executable
	if _, err := exec.LookPath(bin); err == nil {
		return true
	}
	return false
}

func (r *WrongHyphenBeforeSubcommandRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "-", " ", 1)
}
