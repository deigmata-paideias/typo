package common

import (
	"regexp"
	"strings"
)

type RemoveShellPromptLiteralRule struct{}

func (r *RemoveShellPromptLiteralRule) ID() string { return "remove_shell_prompt_literal" }

func (r *RemoveShellPromptLiteralRule) Match(command string, output string) bool {
	// match if output says "$: command not found" AND script starts with "$ "
	// regex: ^[\s]*\$ [\S]+
	match, _ := regexp.MatchString(`^[\s]*\$ [\S]+`, command)
	return match && strings.Contains(output, "$: command not found")
}

func (r *RemoveShellPromptLiteralRule) GetNewCommand(command string, output string) string {
	return strings.TrimLeft(command, "$ ")
}
