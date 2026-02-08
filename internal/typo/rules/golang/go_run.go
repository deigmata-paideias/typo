package golang

import (
	"strings"
)

type GoRunRule struct{}

func (r *GoRunRule) ID() string {
	return "go_run"
}

func (r *GoRunRule) Match(command string, output string) bool {
	// go run foo
	// error: go run: no go files listed (this is from thefuck comment, but let's check strict suffix)
	// The rule simple checks: starts with "go run" and does not end with ".go"
	// However, "go run ." is valid. "go run main.go" is valid.
	// "go run foo" -> might expect "foo.go"

	return strings.HasPrefix(command, "go run ") && !strings.HasSuffix(command, ".go") && !strings.HasSuffix(command, ".")
}

func (r *GoRunRule) GetNewCommand(command string, output string) string {
	return command + ".go"
}
