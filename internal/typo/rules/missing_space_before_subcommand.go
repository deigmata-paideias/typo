package rules

import (
	"os/exec"
	"strings"
)

type MissingSpaceBeforeSubcommandRule struct{}

func (r *MissingSpaceBeforeSubcommandRule) ID() string { return "missing_space_before_subcommand" }

func (r *MissingSpaceBeforeSubcommandRule) Match(command string, output string) bool {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return false
	}
	cmd := parts[0]
	// If command exists, we don't fix it (unless we want to? No, usually valid command is preferred)
	if _, err := exec.LookPath(cmd); err == nil {
		return false
	}

	// Try splitting
	for i := 2; i < len(cmd); i++ { // Ignore 1 char prefixes
		prefix := cmd[:i]
		if _, err := exec.LookPath(prefix); err == nil {
			return true
		}
	}
	return false
}

func (r *MissingSpaceBeforeSubcommandRule) GetNewCommand(command string, output string) string {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return command
	}
	cmd := parts[0]

	for i := 2; i < len(cmd); i++ {
		prefix := cmd[:i]
		if _, err := exec.LookPath(prefix); err == nil {
			// e.g. gitcommit -> git commit
			// Replace first occurrence of 'git' with 'git ' in the command string?
			// Be careful: 'gitcommit' -> 'git commit'
			// strings.Replace(command, cmd, prefix + " " + cmd[i:], 1)
			return strings.Replace(command, cmd, prefix+" "+cmd[i:], 1)
		}
	}
	return command
}
