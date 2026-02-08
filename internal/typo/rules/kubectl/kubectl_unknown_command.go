package kubectl

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type KubectlUnknownCommandRule struct{}

func (r *KubectlUnknownCommandRule) ID() string {
	return "kubectl_unknown_command"
}

func (r *KubectlUnknownCommandRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "kubectl") &&
		(strings.Contains(output, "unknown command") ||
			strings.Contains(output, "Unknown command"))
}

func (r *KubectlUnknownCommandRule) GetNewCommand(command string, output string) string {
	// kubectl: unknown command "stat"
	re := regexp.MustCompile(`unknown command "([^"]+)"`)
	matches := re.FindStringSubmatch(output)

	var brokenCmd string
	if len(matches) > 1 {
		brokenCmd = matches[1]
	}

	if brokenCmd == "" {
		return command
	}

	// Common kubectl commands
	kubectlCmds := []string{
		"get", "describe", "create", "apply", "delete", "edit", "patch", "replace",
		"rollout", "scale", "autoscale", "run", "expose", "set", "label", "annotate",
		"logs", "exec", "attach", "port-forward", "proxy", "cp", "auth", "diff",
		"top", "cordon", "uncordon", "drain", "taint", "cluster-info", "api-resources",
		"api-versions", "config", "plugin", "version", "debug", "events", "wait",
		"kustomize", "certificate", "completion", "convert", "explain", "options",
	}

	bestMatch := utils.Match(brokenCmd, kubectlCmds)
	if bestMatch != "" {
		return strings.Replace(command, brokenCmd, bestMatch, 1)
	}

	return command
}
