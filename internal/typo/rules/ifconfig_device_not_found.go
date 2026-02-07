package rules

import (
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type IfconfigDeviceNotFoundRule struct{}

func (r *IfconfigDeviceNotFoundRule) ID() string { return "ifconfig_device_not_found" }

func (r *IfconfigDeviceNotFoundRule) Match(command string, output string) bool {
	return strings.Contains(command, "ifconfig") && strings.Contains(output, "error fetching interface information: Device not found")
}

func (r *IfconfigDeviceNotFoundRule) GetNewCommand(command string, output string) string {
	// command: ifconfig eth0 ...
	// output: eth0: error ...
	// Extract interface from output?
	// python: interface = command.output.split(' ')[0][:-1]
	// "eth0: error..." -> split " "[0] is "eth0:", [:-1] is "eth0"

	parts := strings.Fields(output)
	if len(parts) == 0 {
		return command
	}
	wrongIface := strings.TrimRight(parts[0], ":")

	// Get available interfaces
	out, err := utils.ExecCommandWithOutput("ifconfig", "-a") // or zsh -c "ifconfig -a"
	if err != nil {
		return command
	}

	var interfaces []string
	// simple parser as per python code: line.split(' ')[0]
	lines := strings.Split(out, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			// In some ifconfig outputs, interface is start of line.
			// python: if line and ... not line.startswith(' '): yield line.split(' ')[0]
			// We should check the original output format.
			// Standard unix ifconfig (BSD/Linux) puts interface at start of line
			if !strings.HasPrefix(line, "\t") && !strings.HasPrefix(line, " ") {
				iface := strings.Fields(line)[0]
				iface = strings.TrimRight(iface, ":")
				interfaces = append(interfaces, iface)
			}
		}
	}

	best := utils.Match(wrongIface, interfaces)
	if best != "" {
		return strings.Replace(command, wrongIface, best, 1)
	}
	return command
}
