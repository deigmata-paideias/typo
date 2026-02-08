package kubectl

import (
	"strings"
)

type KubectlGetMissingResourceRule struct{}

func (r *KubectlGetMissingResourceRule) ID() string {
	return "kubectl_get_missing_resource"
}

func (r *KubectlGetMissingResourceRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "kubectl get") &&
		strings.Contains(output, "You must provide one or more resources")
}

func (r *KubectlGetMissingResourceRule) GetNewCommand(command string, output string) string {
	// Add common resource type suggestion
	suggestions := []string{
		"kubectl get pods",
		"kubectl get services",
		"kubectl get deployments",
		"kubectl get all",
	}

	// Return the first suggestion as the new command
	// Or we could return multiple suggestions
	return suggestions[0]
}
