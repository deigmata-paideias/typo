package common

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type SudoCommandFromUserPathRule struct{}

func (r *SudoCommandFromUserPathRule) ID() string {
	return "sudo_command_from_user_path"
}

func (r *SudoCommandFromUserPathRule) Match(command string, output string) bool {
	// sudo: cmd: command not found
	return strings.HasPrefix(command, "sudo") &&
		strings.Contains(output, "command not found") &&
		r.getCommandName(output) != ""
}

func (r *SudoCommandFromUserPathRule) getCommandName(output string) string {
	re := regexp.MustCompile(`sudo: (.*): command not found`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func (r *SudoCommandFromUserPathRule) GetNewCommand(command string, output string) string {
	cmdName := r.getCommandName(output)
	if cmdName == "" {
		return command
	}

	// Check if cmdName exists in user path (ignoring sudo env)
	// 'which' runs in current process env (user env)
	path, err := utils.ExecCommandWithOutput("which", cmdName)
	if err == nil && path != "" {
		// Suggest: sudo env "PATH=$PATH" cmd ...
		// Replace 'sudo cmd' with 'sudo env "PATH=$PATH" cmd'
		return strings.Replace(command, "sudo "+cmdName, fmt.Sprintf("sudo env \"PATH=$PATH\" %s", cmdName), 1)
	}
	return command
}
