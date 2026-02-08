package hosts

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type HostsCliRule struct{}

func (r *HostsCliRule) ID() string { return "hosts_cli" }

func (r *HostsCliRule) Match(command string, output string) bool {
	return strings.Contains(output, "Error: No such command") ||
		strings.Contains(output, "hostscli.errors.WebsiteImportError")
}

func (r *HostsCliRule) GetNewCommand(command string, output string) string {
	if strings.Contains(output, "hostscli.errors.WebsiteImportError") {
		return "hostscli websites"
	}

	re := regexp.MustCompile(`Error: No such command "(.*)"`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		misspelled := matches[1]
		candidates := []string{"block", "unblock", "websites", "block_all", "unblock_all"}
		best := utils.Match(misspelled, candidates)
		if best != "" {
			return strings.Replace(command, misspelled, best, 1)
		}
	}
	return command
}
