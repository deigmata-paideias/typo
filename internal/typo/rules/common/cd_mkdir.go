package common

import (
	"regexp"
	"strings"
)

type CdMkdirRule struct{}

func (r *CdMkdirRule) ID() string {
	return "cd_mkdir"
}

func (r *CdMkdirRule) Match(command string, output string) bool {
	if !strings.HasPrefix(command, "cd ") {
		return false
	}

	lowerOutput := strings.ToLower(output)
	return strings.Contains(lowerOutput, "no such file or directory") ||
		strings.Contains(lowerOutput, "does not exist") ||
		strings.Contains(lowerOutput, "can't cd to")
}

func (r *CdMkdirRule) GetNewCommand(command string, output string) string {
	// cd foo/bar -> mkdir -p foo/bar && cd foo/bar
	re := regexp.MustCompile(`^cd (.*)`)
	matches := re.FindStringSubmatch(command)
	if len(matches) > 1 {
		dir := matches[1]
		return "mkdir -p " + dir + " && cd " + dir
	}
	return command
}
