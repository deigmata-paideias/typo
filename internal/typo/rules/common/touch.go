package common

import (
	"path/filepath"
	"strings"
)

type TouchRule struct{}

func (r *TouchRule) ID() string {
	return "touch"
}

func (r *TouchRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "touch") &&
		strings.Contains(output, "No such file or directory")
}

func (r *TouchRule) GetNewCommand(command string, output string) string {
	parts := strings.Fields(command)
	if len(parts) < 2 {
		return command
	}

	for _, arg := range parts[1:] {
		dir := filepath.Dir(arg)
		if dir == "." || dir == "/" {
			continue
		}
		return "mkdir -p " + dir + " && " + command
	}

	return command
}
