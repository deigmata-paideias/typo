package rules

import (
	"strings"
)

type YumInvalidOperationRule struct{}

func (r *YumInvalidOperationRule) ID() string { return "yum_invalid_operation" }

func (r *YumInvalidOperationRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "yum") &&
		strings.Contains(output, "Invalid operation")
}

func (r *YumInvalidOperationRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "yum", "yum install", 1)
}
