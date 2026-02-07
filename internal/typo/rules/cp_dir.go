package rules

import (
	"regexp"
	"strings"
)

type CpOmittingDirectoryRule struct{}

func (r *CpOmittingDirectoryRule) ID() string {
	return "cp_omitting_directory"
}

func (r *CpOmittingDirectoryRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "cp") &&
		(strings.Contains(output, "omitting directory") ||
			strings.Contains(output, "is a directory"))
}

func (r *CpOmittingDirectoryRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`^cp `)
	return re.ReplaceAllString(command, "cp -a ")
}
