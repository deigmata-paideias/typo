package tsuru

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type TsuruNotCommandRule struct{}

func (r *TsuruNotCommandRule) ID() string {
	return "tsuru_not_command"
}

func (r *TsuruNotCommandRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "tsuru") &&
		strings.Contains(output, "is not a tsuru command") &&
		strings.Contains(output, "Did you mean?")
}

func (r *TsuruNotCommandRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`tsuru: "([^"]*)" is not a tsuru command`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return command
	}
	broken := matches[1]

	// Extract suggestions
	// Did you mean?
	// 	target-add
	// 	...
	// python code: matches "Did you mean?" line and next lines?
	// get_all_matched_commands(output)

	// We'll parse lines after "Did you mean?"
	var candidates []string
	lines := strings.Split(output, "\n")
	read := false
	for _, line := range lines {
		if strings.Contains(line, "Did you mean?") {
			read = true
			continue
		}
		if read {
			cand := strings.TrimSpace(line)
			if cand != "" {
				candidates = append(candidates, cand)
			}
		}
	}

	best := utils.Match(broken, candidates)
	if best != "" {
		return strings.Replace(command, broken, best, 1)
	}
	return command
}
