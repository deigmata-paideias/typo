package common

import (
	"os"
	"strings"
)

type ChmodXRule struct{}

func (r *ChmodXRule) ID() string {
	return "chmod_x"
}

func (r *ChmodXRule) Match(command string, output string) bool {
	// ./script.sh
	// zsh: permission denied: ./script.sh
	// bash: ./script.sh: Permission denied

	if !strings.HasPrefix(command, "./") && !strings.HasPrefix(command, "/") {
		return false
	}

	if strings.Contains(strings.ToLower(output), "permission denied") {
		// Check if it's a file and not executable
		path := strings.Fields(command)[0]
		info, err := os.Stat(path)
		if err == nil && !info.IsDir() {
			// Check if already executable (skip if it is)
			if info.Mode()&0111 == 0 {
				return true
			}
		}
	}
	return false
}

func (r *ChmodXRule) GetNewCommand(command string, output string) string {
	path := strings.Fields(command)[0]
	// chmod +x ./script.sh && ./script.sh
	return "chmod +x " + path + " && " + command
}
