package mvn

import (
	"strings"
)

type MvnNoCommandRule struct{}

func (r *MvnNoCommandRule) ID() string { return "mvn_no_command" }

func (r *MvnNoCommandRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "mvn") &&
		strings.Contains(output, "No goals have been specified for this build")
}

func (r *MvnNoCommandRule) GetNewCommand(command string, output string) string {
	return command + " clean package"
}
