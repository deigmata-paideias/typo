package mercurial

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type MercurialRule struct{}

func (r *MercurialRule) ID() string { return "mercurial" }

func (r *MercurialRule) Match(command string, output string) bool {
	return strings.Contains(output, "hg: unknown command") &&
		(strings.Contains(output, "(did you mean one of ") || strings.Contains(output, " is ambiguous:"))
}

func (r *MercurialRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`\(did you mean one of ([^\?]+)\?\)`)
	matches := re.FindStringSubmatch(output)
	var possibilities []string
	if len(matches) > 1 {
		possibilities = strings.Split(matches[1], ", ")
	} else {
		// Try ambiguous match regex
		// re.findall(r'\n    ([^$]+)$', command.output)
		// reAmb := regexp.MustCompile("\\n    ([^$]+)$") // This regex seems to match lines indented by 4 spaces at end of string?
		// Go regex: `(?m)^\s{4}(.*)$`
		reAmbGo := regexp.MustCompile(`(?m)^\s{4}(.*)$`)
		ambMatches := reAmbGo.FindAllStringSubmatch(output, -1)
		if len(ambMatches) > 0 && len(ambMatches[0]) > 1 {
			possibilities = strings.Split(ambMatches[0][1], " ") // split the first line of suggestions?
		}
	}

	if len(possibilities) > 0 {
		parts := strings.Fields(command)
		if len(parts) > 1 {
			wrongCmd := parts[1]
			best := utils.Match(wrongCmd, possibilities)
			if best != "" {
				return strings.Replace(command, wrongCmd, best, 1)
			}
		}
	}

	return command
}
