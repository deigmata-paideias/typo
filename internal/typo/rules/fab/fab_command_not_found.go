package fab

import (
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type FabCommandNotFoundRule struct{}

func (r *FabCommandNotFoundRule) ID() string { return "fab_command_not_found" }

func (r *FabCommandNotFoundRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "fab ") &&
		strings.Contains(output, "Warning: Command(s) not found:")
}

func (r *FabCommandNotFoundRule) GetNewCommand(command string, output string) string {
	var notFound []string
	var available []string

	lines := strings.Split(output, "\n")
	state := 0 // 0: init, 1: parsing not found, 2: parsing available

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "Warning: Command(s) not found:") {
			state = 1
			// Extract from same line if present
			parts := strings.SplitN(line, "Warning: Command(s) not found:", 2)
			if len(parts) > 1 && strings.TrimSpace(parts[1]) != "" {
				cmds := strings.Split(parts[1], ",")
				for _, c := range cmds {
					c = strings.TrimSpace(c)
					if c != "" {
						notFound = append(notFound, c)
					}
				}
			}
			continue
		}
		if strings.Contains(line, "Available commands:") {
			state = 2
			continue
		}

		if line == "" {
			continue
		}

		if state == 1 {
			// If we hit available commands, switch state
			if strings.Contains(line, "Available commands:") {
				state = 2
				continue
			}
			// Continuation of not found?
			cmds := strings.Split(line, ",")
			for _, c := range cmds {
				c = strings.TrimSpace(c)
				if c != "" {
					notFound = append(notFound, c)
				}
			}
		} else if state == 2 {
			// Available command
			parts := strings.Fields(line)
			if len(parts) > 0 {
				available = append(available, parts[0])
			}
		}
	}

	newCmd := command
	for _, bad := range notFound {
		if best := utils.Match(bad, available); best != "" {
			newCmd = strings.Replace(newCmd, bad, best, 1)
		}
	}
	return newCmd
}
