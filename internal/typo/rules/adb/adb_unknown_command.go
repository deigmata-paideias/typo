package adb

import (
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type ADBUnknownCommandRule struct{}

func (r *ADBUnknownCommandRule) ID() string {
	return "adb_unknown_command"
}

func (r *ADBUnknownCommandRule) Match(command string, output string) bool {
	// Check if it's adb command
	if !strings.HasPrefix(command, "adb") {
		return false
	}
	// Output starts with "Android Debug Bridge version" means usually bad command or help
	return strings.Contains(output, "Android Debug Bridge version")
}

func (r *ADBUnknownCommandRule) GetNewCommand(command string, output string) string {
	// parse command
	parts := strings.Fields(command)
	if len(parts) < 2 {
		return command
	}

	adbCommands := []string{
		"backup", "bugreport", "connect", "devices", "disable-verity", "disconnect",
		"enable-verity", "emu", "forward", "get-devpath", "get-serialno", "get-state",
		"install", "install-multiple", "jdwp", "keygen", "kill-server", "logcat",
		"pull", "push", "reboot", "reconnect", "restore", "reverse", "root", "run-as",
		"shell", "sideload", "start-server", "sync", "tcpip", "uninstall", "unroot",
		"usb", "wait-for",
	}

	// Iterate parts to find the subcommand.
	// The python rule skips flags like -s, -H, -P, -L and their args.
	// Simplifying: assumes first non-flag arg is the command.

	targetIdx := -1
	for i := 1; i < len(parts); i++ {
		arg := parts[i]
		if strings.HasPrefix(arg, "-") {
			// If it's a flag that takes an argument, skip next.
			// s, H, P, L check
			if arg == "-s" || arg == "-H" || arg == "-P" || arg == "-L" {
				i++ // Skip next arg
			}
			continue
		}
		targetIdx = i
		break
	}

	if targetIdx == -1 {
		return command
	}

	typoCmd := parts[targetIdx]
	closest := utils.Match(typoCmd, adbCommands)
	if closest != "" {
		parts[targetIdx] = closest
		return strings.Join(parts, " ")
	}

	return command
}
