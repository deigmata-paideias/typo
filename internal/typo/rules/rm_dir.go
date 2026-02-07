package rules

import (
	"regexp"
	"strings"
)

type RmDirRule struct{}

func (r *RmDirRule) ID() string {
	return "rm_dir"
}

func (r *RmDirRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "rm") &&
		strings.Contains(strings.ToLower(output), "is a directory")
}

func (r *RmDirRule) GetNewCommand(command string, output string) string {
	// rm dir -> rm -rf dir
	// re.sub('\\brm (.*)', 'rm ' + arguments + ' \\1', command.script)

	// Handles: rm foo
	// output: rm -rf foo

	re := regexp.MustCompile(`\brm (.*)`)
	return re.ReplaceAllString(command, "rm -rf $1")
}
