package common

import (
	"regexp"
	"strings"
)

type LongFormHelpRule struct{}

func (r *LongFormHelpRule) ID() string { return "long_form_help" }

func (r *LongFormHelpRule) Match(command string, output string) bool {
	// Match if output contains help suggestion or mentions --help
	re := regexp.MustCompile(`(?i)(Run|Try) '([^']+)'(?: or '[^']+')? for (?:details|more information)`)
	if re.MatchString(output) {
		return true
	}
	if strings.Contains(output, "--help") {
		return true
	}
	return false
}

func (r *LongFormHelpRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`(?i)(Run|Try) '([^']+)'(?: or '[^']+')? for (?:details|more information)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 2 {
		return matches[2]
	}

	// replace -h with --help
	// only if -h is in the command?
	if strings.Contains(command, " -h") {
		return strings.Replace(command, " -h", " --help", 1)
	}
	// Fallback?
	return command
}
