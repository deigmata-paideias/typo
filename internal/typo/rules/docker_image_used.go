package rules

import (
	"regexp"
	"strings"
)

type DockerImageUsedRule struct{}

func (r *DockerImageUsedRule) ID() string { return "docker_image_used" }

func (r *DockerImageUsedRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "docker") &&
		strings.Contains(output, "image is being used by running container")
}

func (r *DockerImageUsedRule) GetNewCommand(command string, output string) string {
	re := regexp.MustCompile(`image is being used by running container ([a-f0-9]+)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		containerID := matches[1]
		return "docker rm -f " + containerID + " && " + command
	}
	return command
}
