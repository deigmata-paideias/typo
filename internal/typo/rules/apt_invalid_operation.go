package rules

import (
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type AptInvalidOperationRule struct{}

func (r *AptInvalidOperationRule) ID() string { return "apt_invalid_operation" }

func (r *AptInvalidOperationRule) Match(command string, output string) bool {
	return (strings.HasPrefix(command, "apt ") || strings.HasPrefix(command, "apt-get ") || strings.HasPrefix(command, "apt-cache ")) &&
		strings.Contains(output, "E: Invalid operation")
}

func (r *AptInvalidOperationRule) GetNewCommand(command string, output string) string {
	parts := strings.Fields(command)
	if len(parts) < 2 {
		return command
	}

	invalidOp := ""
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, "E: Invalid operation") {
			parts := strings.Split(line, " ")
			if len(parts) > 3 {
				invalidOp = parts[3]
			}
		}
	}

	if invalidOp == "" {
		// Fallback: try to guess from command parts
		// apt-get <op> ...
		invalidOp = parts[1]
	}

	validOps := []string{
		"install", "remove", "purge", "update", "upgrade", "show", "list", "search",
		"autoremove", "check", "clean", "autoclean", "dist-upgrade", "dselect-upgrade",
		"build-dep", "source", "download", "changelog", "moo",
	}

	correction := utils.Match(invalidOp, validOps)
	if correction != "" {
		return strings.Replace(command, invalidOp, correction, 1)
	}

	return command
}
