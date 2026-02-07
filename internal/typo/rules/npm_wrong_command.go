package rules

import (
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type NpmWrongCommandRule struct{}

func (r *NpmWrongCommandRule) ID() string { return "npm_wrong_command" }

func (r *NpmWrongCommandRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "npm") &&
		strings.Contains(output, "where <command> is one of:")
}

func (r *NpmWrongCommandRule) GetNewCommand(command string, output string) string {
	// Parse available commands
	// npm ERR!     access, adduser, audit, ...
	var candidates []string
	lines := strings.Split(output, "\n")
	listing := false
	for _, line := range lines {
		if strings.Contains(line, "where <command> is one of:") {
			listing = true
			continue
		}
		if listing {
			if strings.TrimSpace(line) == "" {
				break
			}
			// npm ERR!     cmd1, cmd2, ...
			// remove npm ERR! prefix if present?
			// Output might just be text depending on npm version.
			// Thefuck splits by comma.
			// Clean line
			clean := strings.TrimSpace(line)
			if strings.HasPrefix(clean, "npm ERR!") {
				clean = strings.TrimPrefix(clean, "npm ERR!")
			}
			parts := strings.Split(clean, ", ")
			for _, part := range parts {
				candidates = append(candidates, strings.TrimSpace(part))
			}
		}
	}

	wrongCmd := ""
	parts := strings.Fields(command)
	for _, p := range parts[1:] {
		if !strings.HasPrefix(p, "-") {
			wrongCmd = p
			break
		}
	}

	if wrongCmd != "" && len(candidates) > 0 {
		best := utils.Match(wrongCmd, candidates)
		if best != "" {
			return strings.Replace(command, wrongCmd, best, 1)
		}
	}

	return command
}
