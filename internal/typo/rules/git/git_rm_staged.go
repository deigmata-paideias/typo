package git

import (
	"strings"
)

type GitRmStagedRule struct{}

func (r *GitRmStagedRule) ID() string { return "git_rm_staged" }

func (r *GitRmStagedRule) Match(command string, output string) bool {
	return strings.Contains(command, "git rm") &&
		strings.Contains(output, "error: the following file has staged content different")
}

func (r *GitRmStagedRule) GetNewCommand(command string, output string) string {
	return command + " --cached"
}
