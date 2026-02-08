package common

import (
	"fmt"
	"regexp"
	"strings"
)

type SshKnownHostsRule struct{}

func (r *SshKnownHostsRule) ID() string {
	return "ssh_known_hosts"
}

func (r *SshKnownHostsRule) Match(command string, output string) bool {
	if !strings.HasPrefix(command, "ssh") && !strings.HasPrefix(command, "scp") {
		return false
	}
	patterns := []string{
		`WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED!`,
		`WARNING: POSSIBLE DNS SPOOFING DETECTED!`,
		`Warning: the \S+ host key for '([^']+)' differs from the key for the IP address '([^']+)'`,
	}
	for _, p := range patterns {
		if regexp.MustCompile(p).MatchString(output) {
			return true
		}
	}
	return false
}

func (r *SshKnownHostsRule) GetNewCommand(command string, output string) string {
	// Pattern: Offending key in /path/to/known_hosts:123
	// Or similar.
	// Python: r'(?:Offending (?:key for IP|\S+ key)|Matching host key) in ([^:]+):(\d+)'
	re := regexp.MustCompile(`(?:Offending (?:key for IP|\S+ key)|Matching host key) in ([^:]+):(\d+)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 2 {
		file := matches[1]
		line := matches[2]
		// sed -i '' '123d' file
		// Note: on linux it is `sed -i` without empty string, on mac it is `sed -i ''`
		// I will generate the mac version since user is on mac.
		// Safe cross-platform way? No easy way.
		return fmt.Sprintf("sed -i '' '%sd' %s && %s", line, file, command)
	}
	return command
}
