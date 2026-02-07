package rules

import (
	"regexp"
	"strings"
)

type NixosCmdNotFoundRule struct{}

func (r *NixosCmdNotFoundRule) ID() string { return "nixos_cmd_not_found" }

func (r *NixosCmdNotFoundRule) Match(command string, output string) bool {
	return strings.Contains(output, "nix-env -iA")
}

func (r *NixosCmdNotFoundRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`nix-env -iA ([^\s]*)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		return "nix-env -iA " + matches[1] + " && " + command
	}
	return command
}
