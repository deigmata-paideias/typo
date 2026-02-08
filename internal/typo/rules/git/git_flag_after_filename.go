package git

import (
	"regexp"
	"strings"
)

type GitFlagAfterFilenameRule struct{}

func (r *GitFlagAfterFilenameRule) ID() string { return "git_flag_after_filename" }

func (r *GitFlagAfterFilenameRule) Match(command string, output string) bool {
	return (strings.Contains(output, "fatal: bad flag '") && strings.Contains(output, "' used after filename")) ||
		(strings.Contains(output, "fatal: option '") && strings.Contains(output, "' must come before non-option arguments"))
}

func (r *GitFlagAfterFilenameRule) GetNewCommand(command string, output string) string {
	re1 := regexp.MustCompile(`fatal: bad flag '(.*?)' used after filename`)
	re2 := regexp.MustCompile(`fatal: option '(.*?)' must come before non-option arguments`)

	match1 := re1.FindStringSubmatch(output)
	var badFlag string
	if len(match1) > 1 {
		badFlag = match1[1]
	} else {
		match2 := re2.FindStringSubmatch(output)
		if len(match2) > 1 {
			badFlag = match2[1]
		}
	}

	if badFlag == "" {
		return command
	}

	parts := strings.Fields(command)
	flagIdx := -1
	for i, p := range parts {
		if p == badFlag {
			flagIdx = i
			break
		}
	}

	if flagIdx == -1 {
		return command
	}

	filenameIdx := -1
	for i := flagIdx - 1; i >= 0; i-- {
		// Heuristic: filename is not "git" and not starting with "-" and usually after subcommand (index > 1)
		// But index 1 is subcommand?
		if !strings.HasPrefix(parts[i], "-") && parts[i] != "git" {
			// Assume parts[0] is git, parts[1] is subcommand (e.g. commit)
			if i > 1 {
				filenameIdx = i
				break
			}
		}
	}

	if filenameIdx != -1 {
		parts[flagIdx], parts[filenameIdx] = parts[filenameIdx], parts[flagIdx]
		return strings.Join(parts, " ")
	}

	return command
}
