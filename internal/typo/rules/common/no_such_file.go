package common

import (
	"regexp"
	"strings"
)

type NoSuchFileRule struct{}

func (r *NoSuchFileRule) ID() string { return "no_such_file" }

func (r *NoSuchFileRule) Match(command string, output string) bool {
	// Patterns:
	// mv: cannot move '...' to '...': No such file or directory
	// cp: cannot create regular file '...': No such file or directory
	// Not a directory
	return (strings.HasPrefix(command, "mv") || strings.HasPrefix(command, "cp")) &&
		(strings.Contains(output, "No such file or directory") || strings.Contains(output, "Not a directory"))
}

func (r *NoSuchFileRule) GetNewCommand(command string, output string) string {
	patterns := []string{
		`mv: cannot move '[^']*' to '([^']*)': No such file or directory`,
		`mv: cannot move '[^']*' to '([^']*)': Not a directory`,
		`cp: cannot create regular file '([^']*)': No such file or directory`,
		`cp: cannot create regular file '([^']*)': Not a directory`,
	}

	for _, p := range patterns {
		re := regexp.MustCompile(p)
		matches := re.FindStringSubmatch(output)
		if len(matches) > 1 {
			file := matches[1]
			// Dir is parent of file
			lastSlash := strings.LastIndex(file, "/")
			if lastSlash != -1 {
				dir := file[:lastSlash]
				return "mkdir -p " + dir + " && " + command
			}
		}
	}
	return command
}
