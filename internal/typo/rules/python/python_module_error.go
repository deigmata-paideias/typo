package python

import (
	"regexp"
	"strings"
)

type PythonModuleErrorRule struct{}

func (r *PythonModuleErrorRule) ID() string { return "python_module_error" }

func (r *PythonModuleErrorRule) Match(command string, output string) bool {
	return strings.Contains(output, "ModuleNotFoundError: No module named '")
}

func (r *PythonModuleErrorRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`ModuleNotFoundError: No module named '([^']+)'`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		module := matches[1]
		return "pip install " + module + " && " + command
	}
	return command
}
