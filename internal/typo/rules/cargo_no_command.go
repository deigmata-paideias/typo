package rules

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type CargoNoCommandRule struct{}

func (r *CargoNoCommandRule) ID() string { return "cargo_no_command" }

func (r *CargoNoCommandRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "cargo") &&
		strings.Contains(output, "no such subcommand")
}

func (r *CargoNoCommandRule) GetNewCommand(command string, output string) string {
	// cargo: 'stat' is not a cargo command.
	re := regexp.MustCompile("no such subcommand: `([^`]+)`")
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return command
	}
	typo := matches[1]

	validCommands := []string{
		"build", "check", "clean", "doc", "new", "init", "run", "test", "bench", "update",
		"search", "publish", "install", "uninstall", "login", "package", "owner", "yank",
		"help", "version",
	}

	best := utils.Match(typo, validCommands)
	if best != "" {
		return strings.Replace(command, typo, best, 1)
	}
	return command
}
