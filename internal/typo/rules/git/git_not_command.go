package git

import (
	"regexp"
	"strings"
)

type GitNotCommandRule struct{}

func (r *GitNotCommandRule) ID() string {
	return "git_not_command"
}

func (r *GitNotCommandRule) Match(command string, output string) bool {
	// "git: 'foo' is not a git command. See 'git --help'."
	// AND "The most similar command is" OR "Did you mean this?"
	return strings.Contains(output, "is not a git command") &&
		(strings.Contains(output, "The most similar command") || strings.Contains(output, "Did you mean"))
}

func (r *GitNotCommandRule) GetNewCommand(command string, output string) string {
	// git: 'foo' is not a git command
	reBroken := regexp.MustCompile("git: '([^']*)' is not a git command")
	matchesBroken := reBroken.FindStringSubmatch(output)
	if len(matchesBroken) < 2 {
		return command
	}
	broken := matchesBroken[1]

	// Parse suggestions
	// The most similar command is\n\tfoo
	// Did you mean this?\n\tfoo

	// reSuggest := regexp.MustCompile("(?:\\n|\\t)\\s*([a-z]+)")
	// Simple regex to catch the command on next line or indented.
	// Python Rule uses `get_all_matched_commands`.

	// Attempt to extract the word after "command is" or "mean this?"
	// Output:
	// ...
	// The most similar command is
	//     status

	var suggestion string
	lines := strings.Split(output, "\n")
	foundMarker := false
	for _, line := range lines {
		if strings.Contains(line, "The most similar command") || strings.Contains(line, "Did you mean") {
			foundMarker = true
			continue
		}
		if foundMarker {
			trimmed := strings.TrimSpace(line)
			if trimmed != "" {
				suggestion = trimmed
				break
			}
		}
	}

	if suggestion != "" {
		// "match-name" might be single word?
		// replace broken with suggestion
		return strings.Replace(command, broken, suggestion, 1)
	}

	return command
}
