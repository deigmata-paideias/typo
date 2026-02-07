package rules

import (
	"strings"
)

type GitPushForceRule struct{}

func (r *GitPushForceRule) ID() string {
	return "git_push_force"
}

func (r *GitPushForceRule) Match(command string, output string) bool {
	// push
	// ! [rejected]
	// failed to push some refs to
	// Updates were rejected because the tip of your current branch is behind
	return strings.Contains(command, "push") &&
		strings.Contains(output, "! [rejected]") &&
		strings.Contains(output, "failed to push some refs to") &&
		strings.Contains(output, "Updates were rejected because the tip of your current branch is behind")
}

func (r *GitPushForceRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "push", "push --force-with-lease", 1)
}
