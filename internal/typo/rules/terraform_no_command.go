package rules

import (
	"regexp"
	"strings"
)

type TerraformNoCommandRule struct{}

func (r *TerraformNoCommandRule) ID() string {
	return "terraform_no_command"
}

func (r *TerraformNoCommandRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "terraform") &&
		strings.Contains(output, "Terraform has no command named \"") &&
		strings.Contains(output, "Did you mean \"")
}

func (r *TerraformNoCommandRule) GetNewCommand(command string, output string) string {
	// Terraform has no command named "pln". Did you mean "plan"?
	reMistake := regexp.MustCompile(`Terraform has no command named "([^"]+)"`)
	reFix := regexp.MustCompile(`Did you mean "([^"]+)"`)

	mistakeMatches := reMistake.FindStringSubmatch(output)
	fixMatches := reFix.FindStringSubmatch(output)

	if len(mistakeMatches) > 1 && len(fixMatches) > 1 {
		return strings.Replace(command, mistakeMatches[1], fixMatches[1], 1)
	}
	return command
}
