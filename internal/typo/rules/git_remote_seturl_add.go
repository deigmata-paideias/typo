package rules

import (
	"strings"
)

type GitRemoteSeturlAddRule struct{}

func (r *GitRemoteSeturlAddRule) ID() string { return "git_remote_seturl_add" }

func (r *GitRemoteSeturlAddRule) Match(command string, output string) bool {
	return strings.Contains(command, "set-url") &&
		strings.Contains(output, "fatal: No such remote")
}

func (r *GitRemoteSeturlAddRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "set-url", "add", 1)
}
