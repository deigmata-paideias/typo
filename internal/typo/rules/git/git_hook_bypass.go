package git

import (
	"strings"
)

type GitHookBypassRule struct{}

func (r *GitHookBypassRule) ID() string {
	return "git_hook_bypass"
}

func (r *GitHookBypassRule) Match(command string, output string) bool {
	// am, commit, push
	parts := strings.Fields(command)
	for _, p := range parts {
		if p == "am" || p == "commit" || p == "push" {
			return true
		}
	}
	return false
}

func (r *GitHookBypassRule) GetNewCommand(command string, output string) string {
	// replace cmd with cmd --no-verify
	parts := strings.Fields(command)
	for _, p := range parts {
		if p == "am" || p == "commit" || p == "push" {
			// insert --no-verify after it
			// Simple replace strings might be risky if "commit" appears in message.
			// But usually "git commit".
			return strings.Replace(command, p, p+" --no-verify", 1)
		}
	}
	return command
}
