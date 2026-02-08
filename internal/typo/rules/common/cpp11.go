package common

import (
	"strings"
)

type Cpp11Rule struct{}

func (r *Cpp11Rule) ID() string {
	return "cpp11"
}

func (r *Cpp11Rule) Match(command string, output string) bool {
	return (strings.HasPrefix(command, "g++") || strings.HasPrefix(command, "clang++")) &&
		(strings.Contains(output, "This file requires compiler and library support for the ISO C++ 2011 standard") ||
			strings.Contains(output, "-Wc++11-extensions"))
}

func (r *Cpp11Rule) GetNewCommand(command string, output string) string {
	return command + " -std=c++11"
}
