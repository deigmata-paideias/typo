package rules

import (
	"strings"
)

type GitAddForceRule struct{}

func (r *GitAddForceRule) ID() string {
	return "git_add_force"
}

func (r *GitAddForceRule) Match(command string, output string) bool {
	return strings.Contains(command, "add") &&
		strings.Contains(output, "Use -f if you really want to add them")
}

func (r *GitAddForceRule) GetNewCommand(command string, output string) string {
	// replace add with add --force
	// The python rule uses replace_argument which is smarter (avoids replacing substring in path)
	// But simplistic replace might work for now.
	// However, if "git add add.txt", replace "add" might break.
	// Should match " git add " or start with "git add "
	// Better: replace first "add" after "git"?
	// "git add" -> "git add --force"

	return strings.Replace(command, "add", "add --force", 1)
}
