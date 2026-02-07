package rules

import (
	"regexp"
	"strings"
)

type GitTwoDashesRule struct{}

func (r *GitTwoDashesRule) ID() string { return "git_two_dashes" }

func (r *GitTwoDashesRule) Match(command string, output string) bool {
	return strings.Contains(output, "error: did you mean `") &&
		strings.Contains(output, "` (with two dashes ?)")
}

func (r *GitTwoDashesRule) GetNewCommand(command string, output string) string {
	// re := regexp.MustCompile("error: did you mean `(-+)` \\(with two dashes \\?\\)")
	// The python code assumes output snippet: `error: did you mean `--flag` (with two dashes ?)`
	// But let's look at the python code: to = variable.split('`')[1]
	// It extracts whatever is in the backticks.
	// If the output is `error: did you mean `--help` (with two dashes ?)`
	// `to` = `--help`.
	// `to[1:]` = `-help`.
	// New command replaces `-help` with `--help`.
	// Wait, if I type `git commit -amend`, git says `error: did you mean '--amend' (with two dashes ?)`.
	// So I replace `-amend` with `--amend`.

	// I will use regex to capture the `TEXT` in `error: did you mean `TEXT` (with two dashes ?)`
	reExtract := regexp.MustCompile("error: did you mean `(.*)` \\(with two dashes \\?\\)")
	matches := reExtract.FindStringSubmatch(output)
	if len(matches) > 1 {
		suggestion := matches[1]
		// suggestion is like "--amend"
		if strings.HasPrefix(suggestion, "--") {
			oldFlag := suggestion[1:] // "-amend"
			return strings.Replace(command, oldFlag, suggestion, 1)
		}
	}
	// Fallback, should not happen if Match assumes correct format
	return command
}
