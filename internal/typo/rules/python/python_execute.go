package python

import (
	"strings"
)

type PythonExecuteRule struct{}

func (r *PythonExecuteRule) ID() string {
	return "python_execute"
}

func (r *PythonExecuteRule) Match(command string, output string) bool {
	// python foo
	// checks if not ends with .py
	if !strings.HasPrefix(command, "python ") {
		return false
	}
	return !strings.HasSuffix(command, ".py")
}

func (r *PythonExecuteRule) GetNewCommand(command string, output string) string {
	return command + ".py"
}
