package rules

import (
	"strings"
)

type CpCreateDestinationRule struct{}

func (r *CpCreateDestinationRule) ID() string {
	return "cp_create_destination"
}

func (r *CpCreateDestinationRule) Match(command string, output string) bool {
	// cp or mv
	if !strings.HasPrefix(command, "cp ") && !strings.HasPrefix(command, "mv ") {
		return false
	}

	return strings.Contains(output, "No such file or directory") ||
		(strings.Contains(output, "cp: directory") && strings.Contains(output, "does not exist"))
}

func (r *CpCreateDestinationRule) GetNewCommand(command string, output string) string {
	// mkdir -p last_arg && command
	parts := strings.Fields(command)
	if len(parts) < 2 {
		return command
	}
	dest := parts[len(parts)-1]

	// logic: ensure dest is a directory?
	// In python rule: mkdir -p dest

	return "mkdir -p " + dest + " && " + command
}
