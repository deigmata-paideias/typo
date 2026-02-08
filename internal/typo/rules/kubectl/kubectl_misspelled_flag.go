package kubectl

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type KubectlMisspelledFlagRule struct{}

func (r *KubectlMisspelledFlagRule) ID() string {
	return "kubectl_misspelled_flag"
}

func (r *KubectlMisspelledFlagRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "kubectl") &&
		(strings.Contains(output, "unknown flag") ||
			strings.Contains(output, "unknown shorthand flag"))
}

func (r *KubectlMisspelledFlagRule) GetNewCommand(command string, output string) string {
	// Error: unknown flag: --namepsace
	re := regexp.MustCompile(`unknown (?:shorthand )?flag: (-+[^\s]+)`)
	matches := re.FindStringSubmatch(output)

	var brokenFlag string
	if len(matches) > 1 {
		brokenFlag = matches[1]
	}

	if brokenFlag == "" {
		return command
	}

	// Common kubectl flags
	commonFlags := []string{
		"--namespace", "-n", "--all-namespaces", "-A",
		"--output", "-o", "--selector", "-l",
		"--filename", "-f", "--recursive", "-R",
		"--watch", "-w", "--dry-run", "--force",
		"--grace-period", "--timeout", "--context",
		"--cluster", "--user", "--kubeconfig",
		"--field-selector", "--show-labels", "--show-kind",
		"--sort-by", "--no-headers", "--chunk-size",
		"--container", "-c", "--previous", "-p",
		"--timestamps", "--tail", "--since", "--since-time",
		"--limit-bytes", "--follow", "-f",
		"--record", "--cascade", "--wait",
		"--prune", "--prune-whitelist", "--validate",
		"--server-side", "--field-manager",
	}

	// Remove leading dashes for matching, but keep track of format
	cleanBroken := strings.TrimLeft(brokenFlag, "-")
	cleanFlags := make([]string, len(commonFlags))
	for i, flag := range commonFlags {
		cleanFlags[i] = strings.TrimLeft(flag, "-")
	}

	bestMatch := utils.Match(cleanBroken, cleanFlags)
	if bestMatch != "" {
		// Restore the dash format from the original
		prefix := strings.TrimSuffix(brokenFlag, cleanBroken)
		// Find the original flag format
		for _, flag := range commonFlags {
			if strings.TrimLeft(flag, "-") == bestMatch {
				return strings.Replace(command, brokenFlag, flag, 1)
			}
		}
		return strings.Replace(command, brokenFlag, prefix+bestMatch, 1)
	}

	return command
}
