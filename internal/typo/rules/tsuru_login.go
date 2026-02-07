package rules

import (
	"strings"
)

type TsuruLoginRule struct{}

func (r *TsuruLoginRule) ID() string {
	return "tsuru_login"
}

func (r *TsuruLoginRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "tsuru") &&
		strings.Contains(output, "not authenticated") &&
		strings.Contains(output, "session has expired")
}

func (r *TsuruLoginRule) GetNewCommand(command string, output string) string {
	return "tsuru login && " + command
}
