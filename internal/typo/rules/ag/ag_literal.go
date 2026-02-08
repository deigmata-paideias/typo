package ag

import (
	"strings"
)

type AgLiteralRule struct{}

func (r *AgLiteralRule) ID() string {
	return "ag_literal"
}

func (r *AgLiteralRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "ag") && strings.Contains(output, "run ag with -Q")
}

func (r *AgLiteralRule) GetNewCommand(command string, output string) string {
	// replace first "ag" with "ag -Q"
	return strings.Replace(command, "ag", "ag -Q", 1)
}
