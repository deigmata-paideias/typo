package rules

import (
	"regexp"
	"strings"

	"github.com/deigmata-paideias/typo/internal/utils"
)

type DockerUnknownCommandRule struct{}

func (r *DockerUnknownCommandRule) ID() string {
	return "docker_unknown_command"
}

func (r *DockerUnknownCommandRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "docker") &&
		(strings.Contains(output, "is not a docker command") ||
			strings.Contains(output, "Usage:	docker"))
}

func (r *DockerUnknownCommandRule) GetNewCommand(command string, output string) string {
	// docker: 'bulid' is not a docker command.
	re := regexp.MustCompile(`docker: '(\w+)' is not a docker command`)
	matches := re.FindStringSubmatch(output)

	var brokenCmd string
	if len(matches) > 1 {
		brokenCmd = matches[1]
	}

	if brokenCmd == "" {
		return command
	}

	// Common docker commands
	dockerCmds := []string{
		"attach", "build", "builder", "checkpoint", "commit", "config", "container",
		"context", "cp", "create", "diff", "events", "exec", "export", "history",
		"image", "images", "import", "info", "inspect", "kill", "load", "login",
		"logout", "logs", "manifest", "network", "node", "pause", "plugin", "port",
		"ps", "pull", "push", "rename", "restart", "rm", "rmi", "run", "save",
		"search", "secret", "service", "stack", "start", "stats", "stop", "swarm",
		"system", "tag", "top", "trust", "unpause", "update", "version", "volume",
		"wait",
	}

	bestMatch := utils.Match(brokenCmd, dockerCmds)
	if bestMatch != "" {
		return strings.Replace(command, brokenCmd, bestMatch, 1)
	}

	return command
}
