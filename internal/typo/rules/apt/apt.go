package apt

import (
	"strings"
)

type AptGetSearchRule struct{}

func (r *AptGetSearchRule) ID() string { return "apt_get_search" }

func (r *AptGetSearchRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "apt-get search")
}

func (r *AptGetSearchRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "apt-get search", "apt-cache search", 1)
}

type AptListUpgradableRule struct{}

func (r *AptListUpgradableRule) ID() string { return "apt_list_upgradable" }

func (r *AptListUpgradableRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "apt") &&
		strings.Contains(output, "apt list --upgradable")
}

func (r *AptListUpgradableRule) GetNewCommand(command string, output string) string {
	return "apt list --upgradable"
}

type AptUpgradeRule struct{}

func (r *AptUpgradeRule) ID() string { return "apt_upgrade" }

func (r *AptUpgradeRule) Match(command string, output string) bool {
	return command == "apt list --upgradable" &&
		strings.Count(output, "\n") > 1
}

func (r *AptUpgradeRule) GetNewCommand(command string, output string) string {
	return "apt upgrade"
}
