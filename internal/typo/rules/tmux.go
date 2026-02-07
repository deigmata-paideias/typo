package rules

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type TmuxRule struct{}

func (r *TmuxRule) ID() string {
	return "tmux"
}

func (r *TmuxRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "tmux") &&
		strings.Contains(output, "ambiguous command:") &&
		strings.Contains(output, "could be:")
}

func (r *TmuxRule) GetNewCommand(command string, output string) string {
	// ambiguous command: l, could be: list-buffers, list-clients, ...
	re := regexp.MustCompile(`ambiguous command: (.*), could be: (.*)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 2 {
		wrong := matches[1]
		suggestions := strings.Split(matches[2], ",")
		var candidates []string
		for _, s := range suggestions {
			candidates = append(candidates, strings.TrimSpace(s))
		}

		best := utils.Match(wrong, candidates) // Actually tmux suggests match, we should maybe offer all or just pick one.
		if best != "" {
			return strings.Replace(command, wrong, best, 1)
		}

		// "could be" returns a list.
		// logic: replace wrong with first suggestion or iterate?
		// python: return replace_command(command, old_cmd, suggestions) -> returns list of commands
		// typo's GetNewCommand returns string.
		// I'll pick the first one for now or loop if I change interface.
		// I'll pick `best` for now if match works, or just candidates[0]
		if len(candidates) > 0 {
			// Try to find the closest match to 'wrong' among candidates, although all are partial matches.
			// Just pick the first one.
			return strings.Replace(command, wrong, candidates[0], 1)
		}
	}
	return command
}
