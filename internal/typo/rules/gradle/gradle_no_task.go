package gradle

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type GradleNoTaskRule struct{}

func (r *GradleNoTaskRule) ID() string { return "gradle_no_task" }

func (r *GradleNoTaskRule) Match(command string, output string) bool {
	return (strings.HasPrefix(command, "gradle") || strings.HasPrefix(command, "./gradlew")) &&
		strings.Contains(output, "Task '") &&
		(strings.Contains(output, "not found") || strings.Contains(output, "is ambiguous"))
}

func (r *GradleNoTaskRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`Task '([^']*)' (?:is ambiguous|not found)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return command
	}
	wrongTask := matches[1]

	// Determine binary: gradle or ./gradlew
	binary := strings.Fields(command)[0]
	out, err := utils.ExecCommandWithOutput("zsh", "-c", binary+" tasks")
	if err != nil {
		return command
	}

	var tasks []string
	lines := strings.Split(out, "\n")
	shouldYield := false
	for _, line := range lines {
		line = strings.TrimRight(line, "\r") // handle windows line endings if any, though Mac
		if strings.HasPrefix(line, "----") {
			shouldYield = true
			continue
		}
		if strings.TrimSpace(line) == "" {
			if shouldYield {
				shouldYield = false
			}
			continue
		}
		if shouldYield && !strings.HasPrefix(line, "All tasks runnable") {
			parts := strings.Fields(line)
			if len(parts) > 0 {
				tasks = append(tasks, parts[0])
			}
		}
	}

	best := utils.Match(wrongTask, tasks)
	if best != "" {
		return strings.Replace(command, wrongTask, best, 1)
	}
	return command
}
