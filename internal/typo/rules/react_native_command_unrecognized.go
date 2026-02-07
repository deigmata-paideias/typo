package rules

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type ReactNativeCommandUnrecognizedRule struct{}

func (r *ReactNativeCommandUnrecognizedRule) ID() string { return "react_native_command_unrecognized" }

func (r *ReactNativeCommandUnrecognizedRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "react-native") &&
		strings.Contains(output, "Unrecognized command '")
}

func (r *ReactNativeCommandUnrecognizedRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`Unrecognized command '([^']*)'`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return command
	}
	wrongCmd := matches[1]

	// Get available commands
	// react-native --help
	out, err := utils.ExecCommandWithOutput("react-native", "--help")
	if err != nil {
		return command
	}

	var commands []string
	// output parser
	// Commands:
	//   start
	//   run-ios
	//   ...
	lines := strings.Split(out, "\n")
	read := false
	for _, line := range lines {
		if strings.Contains(line, "Commands:") {
			read = true
			continue
		}
		if read {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			// Extract command name (first word)
			parts := strings.Fields(line)
			if len(parts) > 0 {
				commands = append(commands, parts[0])
			}
		}
	}

	best := utils.Match(wrongCmd, commands)
	if best != "" {
		return strings.Replace(command, wrongCmd, best, 1)
	}
	return command
}
