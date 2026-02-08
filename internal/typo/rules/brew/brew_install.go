package brew

import (
	"regexp"
	"strings"
)

type BrewInstallRule struct{}

func (r *BrewInstallRule) ID() string {
	return "brew_install"
}

func (r *BrewInstallRule) Match(command string, output string) bool {
	// Check for "No available formula" and "Did you mean"
	return strings.Contains(command, "install") &&
		strings.Contains(output, "No available formula") &&
		strings.Contains(output, "Did you mean")
}

func (r *BrewInstallRule) GetNewCommand(command string, output string) string {
	// Warning: No available formula with the name "(?:[^"]+)". Did you mean (.+)\?
	re := regexp.MustCompile(`Did you mean (.+)\?`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return command
	}

	suggestionsStr := matches[1]
	// "formula1, formula2 or formula3" -> ["formula1", "formula2", "formula3"]
	suggestionsStr = strings.ReplaceAll(suggestionsStr, " or ", ", ")
	suggestions := strings.Split(suggestionsStr, ", ")

	if len(suggestions) > 0 {
		// Return first one or logic to return multiple?
		// Rules interface: GetNewCommand returns string.
		// We'll pick the first one.
		return "brew install " + strings.TrimSpace(suggestions[0])
	}

	return command
}
