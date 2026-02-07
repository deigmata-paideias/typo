package rules

import (
	"strings"
)

type TerraformInitRule struct{}

func (r *TerraformInitRule) ID() string {
	return "terraform_init"
}

func (r *TerraformInitRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "terraform") &&
		(strings.Contains(strings.ToLower(output), "this module is not yet installed") ||
			strings.Contains(strings.ToLower(output), "initialization required"))
}

func (r *TerraformInitRule) GetNewCommand(command string, output string) string {
	return "terraform init && " + command
}
