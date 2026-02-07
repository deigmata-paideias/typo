package rules

import (
	"strings"
)

type PipInstallRule struct{}

func (r *PipInstallRule) ID() string { return "pip_install" }

func (r *PipInstallRule) Match(command string, output string) bool {
	return strings.Contains(command, "pip install") && strings.Contains(output, "Permission denied")
}

func (r *PipInstallRule) GetNewCommand(command string, output string) string {
	if !strings.Contains(command, "--user") {
		return strings.Replace(command, " install ", " install --user ", 1)
	}
	// If --user failed or was already there (though logic above prioritizes --user), try sudo
	// But if --user is there, we should probably strip it if we use sudo?
	// The python code:
	// if '--user' not in script: return script.replace(' install ', ' install --user ')
	// return 'sudo {}'.format(script.replace(' --user', ''))
	cleanCmd := strings.Replace(command, " --user", "", 1)
	return "sudo " + cleanCmd
}
