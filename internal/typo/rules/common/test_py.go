package common

import (
	"strings"
)

type TestPyRule struct{}

func (r *TestPyRule) ID() string {
	return "test_py"
}

func (r *TestPyRule) Match(command string, output string) bool {
	return command == "test.py" && strings.Contains(output, "not found")
}

func (r *TestPyRule) GetNewCommand(command string, output string) string {
	return "pytest"
}
