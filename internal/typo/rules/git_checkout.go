package rules

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type GitCheckoutRule struct{}

func (r *GitCheckoutRule) ID() string {
	return "git_checkout"
}

func (r *GitCheckoutRule) Match(command string, output string) bool {
	// "pathspec '...' did not match any file(s) known to git"
	return strings.Contains(output, "did not match any file(s) known to git") &&
		!strings.Contains(output, "Did you forget to 'git add'?")
}

func (r *GitCheckoutRule) GetNewCommand(command string, output string) string {
	// Two strategies:
	// 1. If it looks like creating a branch (and invalid pathspec), suggest -b
	// 2. If it's a typo of an existing branch, suggest the closest branch

	// Extract the missing file/branch name
	re := regexp.MustCompile(`pathspec '([^']*)' did not match`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return command
	}
	missingName := matches[1]

	// 1. Try to find close matches in existing branches
	branchesStr, err := utils.ExecCommand("git", "branch", "-a")
	if err == nil {
		branches := []string{}
		lines := strings.Split(branchesStr, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			line = strings.TrimPrefix(line, "* ")
			if line == "" || strings.Contains(line, "->") {
				continue
			}
			if strings.HasPrefix(line, "remotes/") {
				parts := strings.Split(line, "/")
				if len(parts) > 2 {
					line = strings.Join(parts[2:], "/")
				}
			}
			branches = append(branches, line)
		}

		closest := utils.Match(missingName, branches)
		// If we found a good match (strsim usually returns something, we should check score if possible but Match doesn't return score)
		// Let's rely on Match behavior for now.
		if closest != "" && closest != missingName {
			return strings.Replace(command, missingName, closest, 1)
		}
	}

	// 2. If no similar branch found (or logic dictates), suggest -b
	// Logic: "git checkout <name>" -> "git checkout -b <name>"
	parts := strings.Fields(command)
	for i, part := range parts {
		if part == "checkout" && i+1 < len(parts) {
			// Check if next part is the missing name
			// Simple replacement
			return strings.Replace(command, "checkout", "checkout -b", 1)
		}
	}

	return command
}
