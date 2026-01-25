package utils

import (
	"os/exec"
	"strings"

	"github.com/deigmata-paideias/typo/internal/types"
)

// Check if command exists in the system
// c onefetch alias
// which c --> c: aliased to onefetch
// command -v c --> alias c=onefetch
func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func Convert(val, source string) ([]types.Command, error) {

	var commands []types.Command

	switch source {
	case "alias":
		lines := strings.Split(val, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || !strings.Contains(line, "=") {
				continue
			}

			parts := strings.SplitN(line, "=", 2)
			if len(parts) != 2 {
				continue
			}

			aliasName := strings.TrimSpace(parts[0])
			aliasValue := strings.Trim(strings.TrimSpace(parts[1]), `"' `)

			if aliasName == "" || aliasValue == "" {
				continue
			}

			command := types.Command{
				Name:        aliasName,
				Type:        string(types.Alias),
				Source:      "alias",
				Description: "Alias for: " + aliasValue,
			}
			commands = append(commands, command)
		}

	case "git":
		lines := strings.Split(val, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || !strings.Contains(line, "alias.") {
				continue
			}

			// Handle format: alias.br=branch
			parts := strings.SplitN(line, "=", 2)
			if len(parts) != 2 {
				continue
			}

			aliasKey := strings.TrimSpace(parts[0])
			aliasValue := strings.Trim(strings.TrimSpace(parts[1]), `"' `)

			if aliasKey == "" || aliasValue == "" {
				continue
			}

			// Extract alias name, remove "alias." prefix
			if !strings.HasPrefix(aliasKey, "alias.") {
				continue
			}
			aliasName := strings.TrimPrefix(aliasKey, "alias.")

			command := types.Command{
				Name:        aliasName,
				Type:        string(types.Alias),
				Source:      "git",
				Description: "Git alias: " + aliasValue,
			}
			commands = append(commands, command)
		}

	case "man":
		lines := strings.Split(val, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			// Find the first - as separator
			dashIndex := strings.Index(line, " - ")
			if dashIndex == -1 {
				continue
			}

			// Extract command name part
			cmdPart := strings.TrimSpace(line[:dashIndex])
			description := strings.TrimSpace(line[dashIndex+3:])

			// Split multiple command names
			cmdNames := strings.Split(cmdPart, ", ")
			for _, rawCmd := range cmdNames {
				rawCmd = strings.TrimSpace(rawCmd)

				// Only process section 1 commands
				if !strings.HasSuffix(rawCmd, "(1)") {
					continue
				}

				// Remove (1) suffix
				cmdName := strings.TrimSuffix(rawCmd, "(1)")
				if cmdName == "" {
					continue
				}

				command := types.Command{
					Name:        cmdName,
					Type:        string(types.Man),
					Source:      "man",
					Description: description,
				}

				// Check if command actually exists
				if commandExists(cmdName) {
					commands = append(commands, command)
				}

				// Handle subcommands, format: command-subcommand, git-branch, curl-config etc.
				if strings.Contains(cmdName, "-") {
					parts := strings.SplitN(cmdName, "-", 2)
					if len(parts) == 2 {
						mainCmd := parts[0]
						subCmd := parts[1]
						if mainCmd != "" && subCmd != "" {
							// Mark as subcommand, will be processed later
							commands = append(commands, types.Command{
								Name:        cmdName, // Full name
								Type:        string(types.Man),
								Source:      "man-subcommand",
								Description: description,
							})
						}
					}
				}
			}
		}
	}

	return commands, nil
}
