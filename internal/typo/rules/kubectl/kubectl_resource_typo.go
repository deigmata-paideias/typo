package kubectl

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type KubectlResourceTypoRule struct{}

func (r *KubectlResourceTypoRule) ID() string {
	return "kubectl_resource_typo"
}

func (r *KubectlResourceTypoRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "kubectl") &&
		(strings.Contains(output, "the server doesn't have a resource type") ||
			strings.Contains(output, "error: the server doesn't have a resource type"))
}

func (r *KubectlResourceTypoRule) GetNewCommand(command string, output string) string {
	// Error: error: the server doesn't have a resource type "deployment"
	re := regexp.MustCompile(`doesn't have a resource type "([^"]+)"`)
	matches := re.FindStringSubmatch(output)

	var brokenResource string
	if len(matches) > 1 {
		brokenResource = matches[1]
	}

	if brokenResource == "" {
		return command
	}

	// Common kubectl resource types (plural forms)
	resources := []string{
		"pods", "services", "deployments", "replicasets", "statefulsets", "daemonsets",
		"jobs", "cronjobs", "configmaps", "secrets", "namespaces", "nodes",
		"persistentvolumes", "persistentvolumeclaims", "storageclasses",
		"ingresses", "networkpolicies", "serviceaccounts", "roles", "rolebindings",
		"clusterroles", "clusterrolebindings", "endpoints", "events", "limitranges",
		"resourcequotas", "horizontalpodautoscalers", "poddisruptionbudgets",
		"priorityclasses", "runtimeclasses", "volumeattachments",
	}

	// Also add short names and singular forms
	allResources := append(resources,
		"po", "svc", "deploy", "rs", "sts", "ds", "cm", "ns", "no", "pv", "pvc",
		"ing", "netpol", "sa", "hpa", "pdb", "ep", "ev", "sc",
		"pod", "service", "deployment", "replicaset", "statefulset", "daemonset",
		"configmap", "secret", "namespace", "node", "persistentvolume",
		"persistentvolumeclaim", "storageclass", "ingress", "networkpolicy",
		"serviceaccount", "role", "rolebinding", "clusterrole", "clusterrolebinding",
		"endpoint", "event", "limitrange", "resourcequota", "horizontalpodautoscaler",
		"poddisruptionbudget", "priorityclass", "runtimeclass", "volumeattachment",
	)

	bestMatch := utils.Match(brokenResource, allResources)
	if bestMatch != "" {
		return strings.Replace(command, brokenResource, bestMatch, 1)
	}

	return command
}
