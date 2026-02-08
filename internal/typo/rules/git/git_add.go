package git

import (
	"os"
	"regexp"
	"strings"
)

type GitAddRule struct{}

func (r *GitAddRule) ID() string {
	return "git_add"
}

func (r *GitAddRule) Match(command string, output string) bool {
	return strings.Contains(output, "did not match any file(s) known to git") &&
		!strings.Contains(output, "Did you forget to 'git add'?")
}

func (r *GitAddRule) GetNewCommand(command string, output string) string {
	// error: pathspec 'untracked.txt' did not match any file(s) known to git
	re := regexp.MustCompile(`pathspec '([^']*)' did not match`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return command
	}
	pathspec := matches[1]

	// Check if file exists on disk
	if _, err := os.Stat(pathspec); err == nil {
		// git add -- <file> && <original_command>
		return "git add -- " + pathspec + " && " + command
	}

	return command
}
