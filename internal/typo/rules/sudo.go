package rules

import (
	"strings"
)

type SudoRule struct{}

func (r *SudoRule) ID() string {
	return "sudo"
}

func (r *SudoRule) Match(command string, output string) bool {
	if strings.HasPrefix(command, "sudo") {
		return false
	}

	lowerOutput := strings.ToLower(output)
	return strings.Contains(lowerOutput, "permission denied") ||
		strings.Contains(lowerOutput, "eacces") ||
		strings.Contains(lowerOutput, "operation not permitted")
}

func (r *SudoRule) GetNewCommand(command string, output string) string {
	return "sudo " + command

}
