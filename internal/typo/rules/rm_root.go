package rules

import (
	"strings"
)

type RmRootRule struct{}

func (r *RmRootRule) ID() string { return "rm_root" }

func (r *RmRootRule) Match(command string, output string) bool {
	// command must contain "rm" and "/"
	// output has "--no-preserve-root"
	parts := strings.Fields(command)
	hasRm := false
	hasRoot := false
	for _, p := range parts {
		if p == "rm" {
			hasRm = true
		}
		if p == "/" {
			hasRoot = true
		}
	}
	return hasRm && hasRoot &&
		strings.Contains(output, "--no-preserve-root") &&
		!strings.Contains(command, "--no-preserve-root")
}

func (r *RmRootRule) GetNewCommand(command string, output string) string {
	return command + " --no-preserve-root"
}
