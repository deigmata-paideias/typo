package rules

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type GruntTaskNotFoundRule struct{}

func (r *GruntTaskNotFoundRule) ID() string { return "grunt_task_not_found" }

func (r *GruntTaskNotFoundRule) Match(command string, output string) bool {
	return strings.Contains(command, "grunt") && strings.Contains(output, "Warning: Task \"") && strings.Contains(output, "\" not found.")
}

func (r *GruntTaskNotFoundRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`Warning: Task "(.*)" not found.`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return command
	}
	misspelledTask := strings.Split(matches[1], ":")[0] // Handle colon if present? python: [0].split(':')[0]

	out, err := utils.ExecCommandWithOutput("zsh", "-c", "grunt --help")
	if err != nil {
		return command
	}

	var tasks []string
	lines := strings.Split(out, "\n")
	shouldYield := false
	for _, line := range lines {
		if strings.Contains(line, "Available tasks") {
			shouldYield = true
			continue
		}
		if shouldYield && strings.TrimSpace(line) == "" {
			break // or continue and set false
		}
		if shouldYield && strings.Contains(line, "  ") {
			parts := strings.Fields(line)
			if len(parts) > 0 {
				tasks = append(tasks, parts[0])
			}
		}
	}

	best := utils.Match(misspelledTask, tasks)
	if best != "" {
		return strings.Replace(command, misspelledTask, best, 1)
	}
	return command
}
