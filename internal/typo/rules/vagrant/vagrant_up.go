package vagrant

import (
	"strings"
)

type VagrantUpRule struct{}

func (r *VagrantUpRule) ID() string { return "vagrant_up" }

func (r *VagrantUpRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "vagrant") &&
		strings.Contains(strings.ToLower(output), "run `vagrant up`")
}

func (r *VagrantUpRule) GetNewCommand(command string, output string) string {
	parts := strings.Fields(command)
	machine := ""
	if len(parts) >= 3 {
		// vagrant ssh machine ...
		// index 2 might be machine name if command is `vagrant ssh machine`
		// But valid commands are `vagrant ssh [name]`.
		// If len >= 3, maybe. Python uses `cmds[2]`.
		machine = parts[2]
	}

	if machine != "" {
		return "vagrant up " + machine + " && " + command
	}
	return "vagrant up && " + command
}
