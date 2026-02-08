package az

import (
	"regexp"
	"strings"
)

type AzCliRule struct{}

func (r *AzCliRule) ID() string { return "az_cli" }

func (r *AzCliRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "az ") &&
		strings.Contains(output, "is not in the") &&
		strings.Contains(output, "command group")
}

func (r *AzCliRule) GetNewCommand(command string, output string) string {
	// 'something' is not in the 'group' command group. See 'az group --help'.
	// The most similar choice to 'something' is:
	// 	suggested

	reMsg := regexp.MustCompile(`'([^']+)' is not in the '.*' command group\.`)
	matchesMsg := reMsg.FindStringSubmatch(output)
	if len(matchesMsg) < 2 {
		return command
	}
	typo := matchesMsg[1]

	reSugg := regexp.MustCompile(`The most similar choice to '.*' is:\s+(.*)`)
	matchesSugg := reSugg.FindStringSubmatch(output)
	if len(matchesSugg) < 2 {
		return command
	}
	suggestion := strings.TrimSpace(matchesSugg[1])

	return strings.Replace(command, typo, suggestion, 1)
}
