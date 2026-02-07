package rules

import (
	"strings"
)

type SedUnterminatedSRule struct{}

func (r *SedUnterminatedSRule) ID() string {
	return "sed_unterminated_s"
}

func (r *SedUnterminatedSRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "sed") &&
		(strings.Contains(output, "unterminated `s' command") || strings.Contains(output, "unterminated substitute in regular expression"))
}

func (r *SedUnterminatedSRule) GetNewCommand(command string, output string) string {
	// sed s/foo/bar -> sed s/foo/bar/
	parts := strings.Fields(command)

	for i, part := range parts {
		// Clean quotes for checking, but we need to append to the original part
		// This is a naive implementation.
		// If part is 's/foo/bar', we want 's/foo/bar/'

		// Check if it starts with s/ and doesn't end with /
		// Be careful with quotes: 's/foo/bar' ends with ' not /

		cleanPart := strings.Trim(part, "'\"")
		if strings.HasPrefix(cleanPart, "s/") && !strings.HasSuffix(cleanPart, "/") {
			// Append / before the closing quote if it exists
			if strings.HasSuffix(part, "'") {
				parts[i] = part[:len(part)-1] + "/'"
			} else if strings.HasSuffix(part, "\"") {
				parts[i] = part[:len(part)-1] + "/\""
			} else {
				parts[i] = part + "/"
			}
		}
	}

	return strings.Join(parts, " ")
}
