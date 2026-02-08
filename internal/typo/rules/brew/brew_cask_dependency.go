package brew

import (
	"strings"
)

type BrewCaskDependencyRule struct{}

func (r *BrewCaskDependencyRule) ID() string {
	return "brew_cask_dependency"
}

func (r *BrewCaskDependencyRule) Match(command string, output string) bool {
	// install in command AND "brew cask install" in output
	return strings.Contains(command, "install") && strings.Contains(output, "brew cask install")
}

func (r *BrewCaskDependencyRule) GetNewCommand(command string, output string) string {
	// Extract "brew cask install ..." lines from output.
	// Python code handles multiple lines.
	// For simplicity, we can just grab the lines starting with "brew cask install"

	var installCmds []string
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "brew cask install") {
			installCmds = append(installCmds, strings.TrimSpace(line))
		}
	}

	if len(installCmds) > 0 {
		// chained: install_deps && original_command
		return strings.Join(installCmds, " && ") + " && " + command
	}
	return command
}
