package rules

import (
	"regexp"
	"strings"
)

type BrewReinstallRule struct{}

func (r *BrewReinstallRule) ID() string { return "brew_reinstall" }

func (r *BrewReinstallRule) Match(command string, output string) bool {
	return strings.Contains(command, "install") &&
		strings.Contains(output, "is already installed and up-to-date") &&
		strings.Contains(output, "run `brew reinstall")
}

func (r *BrewReinstallRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "install", "reinstall", 1)
}

type BrewUninstallRule struct{}

func (r *BrewUninstallRule) ID() string { return "brew_uninstall" }

func (r *BrewUninstallRule) Match(command string, output string) bool {
	return strings.Contains(command, "brew") &&
		(strings.Contains(command, "uninstall") || strings.Contains(command, "remove") || strings.Contains(command, "rm")) &&
		strings.Contains(output, "brew uninstall --force")
}

func (r *BrewUninstallRule) GetNewCommand(command string, output string) string {
	// brew uninstall package -> brew uninstall --force package
	// Or just insert --force after uninstall
	re := regexp.MustCompile(`(uninstall|remove|rm)`)
	return re.ReplaceAllString(command, "$1 --force")
}

type BrewUpdateFormulaRule struct{}

func (r *BrewUpdateFormulaRule) ID() string { return "brew_update_formula" }

func (r *BrewUpdateFormulaRule) Match(command string, output string) bool {
	return strings.Contains(command, "update") &&
		strings.Contains(output, "Error: This command updates brew itself") &&
		strings.Contains(output, "Use `brew upgrade")
}

func (r *BrewUpdateFormulaRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "update", "upgrade", 1)
}
