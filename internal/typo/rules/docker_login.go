package rules

import (
    "strings"
)

type DockerLoginRule struct{}

func (r *DockerLoginRule) ID() string {
    return "docker_login"
}

func (r *DockerLoginRule) Match(command string, output string) bool {
    return strings.Contains(command, "docker") &&
           strings.Contains(output, "access denied") &&
           strings.Contains(output, "may require 'docker login'")
}

func (r *DockerLoginRule) GetNewCommand(command string, output string) string {
    return "docker login && " + command
}
