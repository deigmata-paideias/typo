package rules

import (
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type PacmanNotFoundRule struct{}

func (r *PacmanNotFoundRule) ID() string { return "pacman_not_found" }

func (r *PacmanNotFoundRule) Match(command string, output string) bool {
	return (strings.HasPrefix(command, "pacman") || strings.HasPrefix(command, "yay") ||
		strings.HasPrefix(command, "yaourt") || strings.HasPrefix(command, "sudo pacman")) &&
		strings.Contains(output, "error: target not found:")
}

func (r *PacmanNotFoundRule) GetNewCommand(command string, output string) string {
	// command: yay -S llc
	// output: error: target not found: llc
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return command
	}
	target := parts[len(parts)-1] // assumption: target is last argument

	// pkgfile <target>
	out, err := utils.ExecCommandWithOutput("pkgfile", target)
	if err != nil || out == "" {
		return command
	}
	// pkgfile llc -> extra/llvm
	// We want "llvm"
	lines := strings.Split(out, "\n")
	if len(lines) > 0 {
		line := lines[0]
		if idx := strings.Index(line, "/"); idx != -1 {
			pkg := strings.TrimSpace(line[idx+1:])
			return strings.Replace(command, target, pkg, 1)
		}
	}
	return command
}
