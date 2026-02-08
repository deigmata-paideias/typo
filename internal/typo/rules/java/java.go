package java

import (
	"strings"
)

type JavaRule struct{}

func (r *JavaRule) ID() string { return "java" }

func (r *JavaRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "java ") && strings.HasSuffix(command, ".java")
}

func (r *JavaRule) GetNewCommand(command string, output string) string {
	return strings.TrimSuffix(command, ".java")
}
