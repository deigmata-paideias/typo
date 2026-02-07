package rules

import (
	"strings"
)

type PhpSRule struct{}

func (r *PhpSRule) ID() string { return "php_s" }

func (r *PhpSRule) Match(command string, output string) bool {
	// match if command has -s but not at the end (which might be valid? python code: last part != -s)
	// command.script_parts[-1] != '-s'
	parts := strings.Fields(command)
	if len(parts) > 1 && parts[0] == "php" {
		hasLowerS := false
		for _, p := range parts {
			if p == "-s" {
				hasLowerS = true
				break
			}
		}
		if hasLowerS && parts[len(parts)-1] != "-s" {
			return true
		}
	}
	return false
}

func (r *PhpSRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "-s", "-S", 1)
}
