package dnf

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type DnfNoSuchCommandRule struct{}

func (r *DnfNoSuchCommandRule) ID() string { return "dnf_no_such_command" }

func (r *DnfNoSuchCommandRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "dnf") &&
		strings.Contains(output, "No such command:")
}

func (r *DnfNoSuchCommandRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`No such command: ([^.]+)\.`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return command
	}
	typo := strings.TrimSpace(matches[1])

	validCommands := []string{
		"alias", "autoremove", "check", "check-update", "clean", "deplist", "distro-sync",
		"downgrade", "group", "help", "history", "info", "install", "list", "makecache",
		"mark", "module", "provides", "reinstall", "remove", "repolist", "repoquery",
		"repository-packages", "search", "shell", "swap", "updateinfo", "upgrade",
		"upgrade-minimal",
	}

	best := utils.Match(typo, validCommands)
	if best != "" {
		return strings.Replace(command, typo, best, 1)
	}
	return command
}
