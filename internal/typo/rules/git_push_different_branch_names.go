package rules

import (
    "regexp"
	"strings"
)

type GitPushDifferentBranchNamesRule struct{}

func (r *GitPushDifferentBranchNamesRule) ID() string {
	return "git_push_different_branch_names"
}

func (r *GitPushDifferentBranchNamesRule) Match(command string, output string) bool {
    return strings.Contains(command, "push") &&
           strings.Contains(output, "The upstream branch of your current branch does not match")
}

func (r *GitPushDifferentBranchNamesRule) GetNewCommand(command string, output string) string {
    // To push to the upstream branch on the remote, use
    //     git push origin HEAD:main
    // Or similar.

    // Regex: ^ +(git push [^\s]+ [^\s]+)  (multiline)
    re := regexp.MustCompile(`(?m)^\s+(git push [^\s]+ [^\s]+)`)
    matches := re.FindStringSubmatch(output)
    if len(matches) > 1 {
        return matches[1]
    }
    return command
}
