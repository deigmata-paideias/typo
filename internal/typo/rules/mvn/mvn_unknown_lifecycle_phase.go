package mvn

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type MvnUnknownLifecyclePhaseRule struct{}

func (r *MvnUnknownLifecyclePhaseRule) ID() string { return "mvn_unknown_lifecycle_phase" }

func (r *MvnUnknownLifecyclePhaseRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "mvn") &&
		strings.Contains(output, "Unknown lifecycle phase")
}

func (r *MvnUnknownLifecyclePhaseRule) GetNewCommand(command string, output string) string {
	// [ERROR] Unknown lifecycle phase "foo". You must specify ...
	// Available lifecycle phases are: validate, initialize, ... -> [Help 1]
	reFailed := regexp.MustCompile(`Unknown lifecycle phase "(.+)"`)
	reAvailable := regexp.MustCompile(`Available lifecycle phases are: (.+) -> \[Help 1\]`)

	failedMatches := reFailed.FindStringSubmatch(output)
	availableMatches := reAvailable.FindStringSubmatch(output)

	if len(failedMatches) > 1 && len(availableMatches) > 1 {
		failed := failedMatches[1]
		availableStr := availableMatches[1]
		available := strings.Split(availableStr, ", ")

		best := utils.Match(failed, available)
		if best != "" {
			return strings.Replace(command, failed, best, 1)
		}
	}

	return command
}
