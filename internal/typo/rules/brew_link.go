package rules

import (
	"strings"
)

type BrewLinkRule struct{}

func (r *BrewLinkRule) ID() string { return "brew_link" }

func (r *BrewLinkRule) Match(command string, output string) bool {
	parts := strings.Fields(command)
	if len(parts) < 2 {
		return false
	}
	isLink := parts[1] == "link" || parts[1] == "ln"
	return isLink && strings.Contains(output, "brew link --overwrite --dry-run")
}

func (r *BrewLinkRule) GetNewCommand(command string, output string) string {
	parts := strings.Fields(command)
	idx := -1
	for i, p := range parts {
		if p == "link" || p == "ln" {
			idx = i
			break
		}
	}
	if idx != -1 {
		parts[idx] = "link"
		// insert --overwrite --dry-run
		rest := append([]string{"--overwrite", "--dry-run"}, parts[idx+1:]...)
		newParts := append(parts[:idx+1], rest...)
		return strings.Join(newParts, " ")
	}
	return command
}
