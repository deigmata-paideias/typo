package common

import (
	"os"
	"strings"
)

type ProveRecursivelyRule struct{}

func (r *ProveRecursivelyRule) ID() string { return "prove_recursively" }

func (r *ProveRecursivelyRule) Match(command string, output string) bool {
	if !strings.HasPrefix(command, "prove") || !strings.Contains(output, "NOTESTS") {
		return false
	}
	// Check if already recursive
	if strings.Contains(command, "-r") || strings.Contains(command, "--recurse") {
		return false
	}
	// Check if any argument is a directory
	parts := strings.Fields(command)
	for _, part := range parts[1:] {
		if !strings.HasPrefix(part, "-") {
			info, err := os.Stat(part)
			if err == nil && info.IsDir() {
				return true
			}
		}
	}
	return false
}

func (r *ProveRecursivelyRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "prove", "prove -r", 1)
}
