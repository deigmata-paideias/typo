package rules

import (
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type PacmanRule struct{}

func (r *PacmanRule) ID() string { return "pacman" }

func (r *PacmanRule) Match(command string, output string) bool {
	// match if command not found
	return strings.Contains(output, "not found") || strings.Contains(output, "command not found")
}

func (r *PacmanRule) GetNewCommand(command string, output string) string {
	// Requires pkgfile
	// pkgfile -b -v <command>
	// command is the script name, e.g. "vim"
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return command
	}
	prog := parts[0]

	out, err := utils.ExecCommandWithOutput("pkgfile", "-b", "-v", prog)
	if err != nil || out == "" {
		return command
	}

	// Output: core/vim 8.2.1-1
	// We want package name: "vim"
	// format: repository/package version
	lines := strings.Split(out, "\n")
	if len(lines) > 0 {
		line := lines[0]
		// split by slash and space
		// core/vim -> vim
		if idx := strings.Index(line, "/"); idx != -1 {
			rest := line[idx+1:]
			if spaceIdx := strings.Index(rest, " "); spaceIdx != -1 {
				pkg := rest[:spaceIdx]
				return "pacman -S " + pkg + " && " + command
			}
		}
	}
	return command
}
