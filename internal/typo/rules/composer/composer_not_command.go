package composer

import (
	"regexp"
	"strings"
)

type ComposerNotCommandRule struct{}

func (r *ComposerNotCommandRule) ID() string { return "composer_not_command" }

func (r *ComposerNotCommandRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "composer") &&
		(strings.Contains(output, "Command") && strings.Contains(output, "is not defined"))
}

func (r *ComposerNotCommandRule) GetNewCommand(command string, output string) string {
	// Command "instlal" is not defined.
	re := regexp.MustCompile(`Command "([^"]+)" is not defined`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return command
	}
	typo := matches[1]

	// There is often a "Did you mean this?" section
	reDidYouMean := regexp.MustCompile(`Did you mean this\?\s+([^\s]+)`)
	matchesDidYouMean := reDidYouMean.FindStringSubmatch(output)
	if len(matchesDidYouMean) > 1 {
		return strings.Replace(command, typo, matchesDidYouMean[1], 1)
	}

	// Fallback to internal list if needed, or rely on output parsing
	// but output parsing is safer if provided
	return command
}
