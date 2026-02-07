package rules

import (
	"strings"
)

type CatDirRule struct{}

func (r *CatDirRule) ID() string {
	return "cat_dir"
}

func (r *CatDirRule) Match(command string, output string) bool {
	// cat: directory_name: Is a directory
	// Mac: cat: directory_name: Is a directory
	return strings.HasPrefix(command, "cat") &&
		strings.Contains(output, "Is a directory")
}

func (r *CatDirRule) GetNewCommand(command string, output string) string {
	return strings.Replace(command, "cat", "ls", 1)
}
