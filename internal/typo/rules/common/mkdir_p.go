package common

import (
	"regexp"
	"strings"
)

type MkdirPRule struct{}

func (r *MkdirPRule) ID() string {
	return "mkdir_p"
}

func (r *MkdirPRule) Match(command string, output string) bool {
	return strings.Contains(command, "mkdir") &&
		strings.Contains(output, "No such file or directory")
}

func (r *MkdirPRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`\bmkdir (.*)`)
	return re.ReplaceAllString(command, "mkdir -p $1")
}
