package rules

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type GitBisectUsageRule struct{}

func (r *GitBisectUsageRule) ID() string {
	return "git_bisect_usage"
}

func (r *GitBisectUsageRule) Match(command string, output string) bool {
	return strings.Contains(command, "bisect") &&
		strings.Contains(output, "usage: git bisect")
}

func (r *GitBisectUsageRule) GetNewCommand(command string, output string) string {
	// usage: git bisect [help|start|bad|good|new|old|terms|skip|next|reset|visualize|replay|log|run]
	// broken command: git bisect strt -> git bisect start

	reBroken := regexp.MustCompile(`git bisect ([^ $]*).*`)
	matchesBroken := reBroken.FindStringSubmatch(command)
	if len(matchesBroken) < 2 {
		return command
	}
	broken := matchesBroken[1]

	reUsage := regexp.MustCompile(`usage: git bisect \[([^\]]+)\]`)
	matchesUsage := reUsage.FindStringSubmatch(output)
	if len(matchesUsage) < 2 {
		return command
	}

	// suggestions separated by |
	suggestions := strings.Split(matchesUsage[1], "|")

	closest := utils.Match(broken, suggestions)
	if closest != "" {
		return strings.Replace(command, broken, closest, 1)
	}

	return command
}
