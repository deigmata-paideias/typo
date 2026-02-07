package rules

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type YarnCommandNotFoundRule struct{}

func (r *YarnCommandNotFoundRule) ID() string { return "yarn_command_not_found" }

func (r *YarnCommandNotFoundRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "yarn") &&
		strings.Contains(output, "error Command \"") &&
		strings.Contains(output, "\" not found.")
}

func (r *YarnCommandNotFoundRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`error Command "(.*)" not found.`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return command
	}
	misspelled := matches[1]

	// NPM commands map
	npmCommands := map[string]string{
		"require": "add",
		"install": "add", // yarn add, not install pkg
	}
	if fix, ok := npmCommands[misspelled]; ok {
		return strings.Replace(command, misspelled, fix, 1)
	}

	// Fetch tasks
	out, err := utils.ExecCommandWithOutput("yarn", "--help")
	if err != nil {
		return command
	}
	var tasks []string
	lines := strings.Split(out, "\n")
	read := false
	for _, line := range lines {
		if strings.Contains(line, "Commands:") {
			read = true
			continue
		}
		if read && strings.Contains(line, "- ") {
			// - command
			parts := strings.Fields(line)
			if len(parts) > 1 {
				tasks = append(tasks, parts[len(parts)-1])
			}
		}
	}

	best := utils.Match(misspelled, tasks)
	if best != "" {
		return strings.Replace(command, misspelled, best, 1)
	}
	return command
}
