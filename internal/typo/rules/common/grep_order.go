package common

import (
	"os"
	"strings"
)

type GrepArgumentsOrderRule struct{}

func (r *GrepArgumentsOrderRule) ID() string {
	return "grep_arguments_order"
}

func (r *GrepArgumentsOrderRule) Match(command string, output string) bool {
	if !strings.HasPrefix(strings.TrimSpace(command), "grep") {
		return false
	}
	if !strings.Contains(output, "No such file or directory") {
		return false
	}

	parts := strings.Fields(command)
	if len(parts) < 3 {
		return false
	}

	for _, part := range parts[1:] {
		if isFile(part) {
			return true
		}
	}

	return false
}

func (r *GrepArgumentsOrderRule) GetNewCommand(command string, output string) string {
	parts := strings.Fields(command)
	var fileArg string
	var fileIndex int

	// Find the actual file argument
	for i, part := range parts {
		if i == 0 {
			continue
		} // skip grep
		if isFile(part) {
			fileArg = part
			fileIndex = i
			break
		}
	}

	if fileArg == "" {
		return command
	}

	// Move file to the end
	newParts := make([]string, 0, len(parts))
	for i, part := range parts {
		if i == fileIndex {
			continue
		}
		newParts = append(newParts, part)
	}
	newParts = append(newParts, fileArg)

	return strings.Join(newParts, " ")
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
