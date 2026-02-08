package kubectl

import (
	"strings"
)

type KubectlNoContextRule struct{}

func (r *KubectlNoContextRule) ID() string {
	return "kubectl_no_context"
}

func (r *KubectlNoContextRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "kubectl") &&
		(strings.Contains(output, "The connection to the server") &&
			strings.Contains(output, "was refused") ||
			strings.Contains(output, "current-context is not set"))
}

func (r *KubectlNoContextRule) GetNewCommand(command string, output string) string {
	// Suggest setting context or viewing available contexts
	if strings.Contains(output, "current-context is not set") {
		return "kubectl config get-contexts"
	}
	return "kubectl config use-context <context-name>"
}
