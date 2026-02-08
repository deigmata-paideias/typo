package git

import (
	"regexp"
	"strings"
)

type GitLfsMistypeRule struct{}

func (r *GitLfsMistypeRule) ID() string { return "git_lfs_mistype" }

func (r *GitLfsMistypeRule) Match(command string, output string) bool {
	return strings.Contains(command, "lfs") && strings.Contains(output, "Did you mean this?")
}

func (r *GitLfsMistypeRule) GetNewCommand(command string, output string) string {
	// Error: unknown command "foo" for "git-lfs"
	reBroken := regexp.MustCompile(`Error: unknown command "([^"]*)" for "git-lfs"`)
	matchesBroken := reBroken.FindStringSubmatch(output)
	if len(matchesBroken) < 2 {
		return command
	}
	broken := matchesBroken[1]

	// Did you mean this?\n\t<suggestion>
	reSuggestion := regexp.MustCompile(`Did you mean this\?\n\t([^\n]+)`)
	matchesSuggestion := reSuggestion.FindStringSubmatch(output)
	if len(matchesSuggestion) < 2 {
		return command
	}
	suggestion := strings.TrimSpace(matchesSuggestion[1])

	return strings.Replace(command, broken, suggestion, 1)
}
