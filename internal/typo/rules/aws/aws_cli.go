package aws

import (
	"regexp"
	"strings"
)

type AwsCliRule struct{}

func (r *AwsCliRule) ID() string {
	return "aws_cli"
}

func (r *AwsCliRule) Match(command string, output string) bool {
	return strings.Contains(output, "usage:") && strings.Contains(output, "maybe you meant:")
}

func (r *AwsCliRule) GetNewCommand(command string, output string) string {
	// Invalid choice: 'foo', maybe you meant:
	// * bar
	// * baz

	// Regex: Invalid choice: '(.*)', maybe you meant:
	reMistake := regexp.MustCompile(`Invalid choice: '(.*)', maybe you meant:`)
	matches := reMistake.FindStringSubmatch(output)
	if len(matches) < 2 {
		return command
	}
	mistake := matches[1]

	// Find all options: * option
	reOptions := regexp.MustCompile(`\*\s(.*)`)
	optionMatches := reOptions.FindAllStringSubmatch(output, -1)

	// Return first for now, or we could return multiple if Rule supported it better.
	// The current interface MatchResult allows multiple candidates, but GetNewCommand returns string.
	// Wait, typo interface returns `string, []MatchResult`.
	// But `Rule` interface in `typo.go`? I need to check `Rule` interface definition.
	// I recall `GetNewCommand` returns `string`.
	// If I want multiple, I might need to change Rule interface or just return best one.
	// Thefuck returns list.

	if len(optionMatches) > 0 && len(optionMatches[0]) > 1 {
		best := optionMatches[0][1]
		// Remove whitespace
		best = strings.TrimSpace(best)
		// Just replace one occurrence
		return strings.Replace(command, mistake, best, 1)
	}

	return command
}
