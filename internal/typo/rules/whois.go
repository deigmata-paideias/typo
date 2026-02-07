package rules

import (
	"net/url"
	"strings"
)

type WhoisRule struct{}

func (r *WhoisRule) ID() string { return "whois" }

func (r *WhoisRule) Match(command string, output string) bool {
	// match any whois command, output not strictly checked in python, just "True" if whois.
	// But `typo` passes output only if "safe".
	return strings.HasPrefix(command, "whois ")
}

func (r *WhoisRule) GetNewCommand(command string, output string) string {
	parts := strings.Fields(command)
	if len(parts) < 2 {
		return command
	}
	arg := parts[1]

	// strip http:// etc
	if strings.Contains(arg, "/") {
		u, err := url.Parse(arg)
		if err == nil && u.Host != "" {
			return strings.Replace(command, arg, u.Host, 1)
		}
		// If parse fails (e.g. no scheme), url key logic might differ.
		// Python uses `urlparse`.
	} else if strings.Count(arg, ".") > 1 {
		// remove left-most subdomain
		// en.wikipedia.org -> wikipedia.org
		dotIdx := strings.Index(arg, ".")
		if dotIdx != -1 {
			newArg := arg[dotIdx+1:]
			return strings.Replace(command, arg, newArg, 1)
		}
	}
	return command
}
