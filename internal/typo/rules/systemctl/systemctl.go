package systemctl

import (
	"strings"
)

type SystemctlRule struct{}

func (r *SystemctlRule) ID() string {
	return "systemctl"
}

func (r *SystemctlRule) Match(command string, output string) bool {
	// systemctl service start -> Unknown operation 'service'.
	// args order confusion
	parts := strings.Fields(command)
	return strings.HasPrefix(command, "systemctl") &&
		strings.Contains(output, "Unknown operation '") &&
		len(parts) >= 3 // systemctl service op ...
}

func (r *SystemctlRule) GetNewCommand(command string, output string) string {
	// Swap last two? systemctl httpd start -> systemctl start httpd
	// Python: cmd[-1], cmd[-2] = cmd[-2], cmd[-1]
	// But it checks index: `len(cmd) - cmd.index('systemctl') == 3`
	// Which means exactly 3 tokens starting from systemctl?
	// `systemctl service start` -> 3 tokens.
	parts := strings.Fields(command)
	// Find "systemctl" index
	idx := -1
	for i, p := range parts {
		if p == "systemctl" {
			idx = i
			break
		}
	}
	if idx != -1 && len(parts)-idx == 3 {
		// Swap last two
		last := parts[len(parts)-1]
		prev := parts[len(parts)-2]
		parts[len(parts)-1] = prev
		parts[len(parts)-2] = last
		return strings.Join(parts, " ")
	}
	return command
}
