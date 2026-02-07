package rules

import (
	// "github.com/xrash/smetrics" // Assuming we can add this dependency or implement Levenshtein calculation
	"strings"
	// But since I cannot easily add dependencies without `go get`, I should check if `match.go` has it.
	// If not, I'll implement a simple one.

	"github.com/deigmata-paideias/typo/internal/utils"
)

type GoUnknownCommandRule struct{}

func (r *GoUnknownCommandRule) ID() string {
	return "go_unknown_command"
}

func (r *GoUnknownCommandRule) Match(command string, output string) bool {
	// Check if it's a go command
	if !strings.HasPrefix(command, "go ") {
		return false
	}
	return strings.Contains(output, "unknown command")
}

func (r *GoUnknownCommandRule) GetNewCommand(command string, output string) string {
	// go unkown -> go known
	parts := strings.Fields(command)
	if len(parts) < 2 {
		return command
	}

	badCmd := parts[1]

	// Standard Go commands
	goCommands := []string{
		"bug", "build", "clean", "doc", "env", "fix", "fmt", "generate",
		"get", "install", "list", "mod", "run", "test", "tool", "version", "vet", "work",
	}

	closest := utils.Match(badCmd, goCommands)
	if closest != "" {
		parts[1] = closest
		return strings.Join(parts, " ")
	}
	return command
}
