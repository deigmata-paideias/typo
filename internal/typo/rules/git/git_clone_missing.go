package git

import (
	"net/url"
	"strings"
)

type GitCloneMissingRule struct{}

func (r *GitCloneMissingRule) ID() string {
	return "git_clone_missing"
}

func (r *GitCloneMissingRule) Match(command string, output string) bool {
	parts := strings.Fields(command)
	if len(parts) != 1 {
		return false
	}

	// Check output for "command not found" etc.
	// Simplifying: if output contains "No such file" or "not found"
	if !strings.Contains(output, "No such file or directory") &&
		!strings.Contains(output, "not found") &&
		!strings.Contains(output, "is not recognised as") {
		return false
	}

	rawURL := parts[0]

	// Check if valid URL
	u, err := url.Parse(rawURL)
	if err == nil && (u.Scheme == "http" || u.Scheme == "https") && u.Host != "" {
		return true
	}

	// Check SSH: git@github.com:user/repo
	if strings.HasPrefix(rawURL, "git@") && strings.Contains(rawURL, ":") {
		return true
	}

	if strings.HasPrefix(rawURL, "ssh://") {
		return true
	}

	return false
}

func (r *GitCloneMissingRule) GetNewCommand(command string, output string) string {
	return "git clone " + command
}
