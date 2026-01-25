package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ExecCommand handles regular cmd processing ("git", "config", "--list")
func ExecCommand(cmd string, args ...string) (string, error) {

	output, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, string(output))
	}

	return string(output), nil
}

// ExecPipeCommand supports pipe commands ExecPipeCommand("git config --list | grep alias")
func ExecPipeCommand(fullCmd string) (string, error) {

	parts := strings.Split(fullCmd, "|")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid pipe command")
	}

	var cmds []*exec.Cmd
	for _, part := range parts {
		part = strings.TrimSpace(part)
		fields := strings.Fields(part)
		if len(fields) == 0 {
			continue
		}
		cmds = append(cmds, exec.Command(fields[0], fields[1:]...))
	}

	if len(cmds) == 0 {
		return "", fmt.Errorf("no valid commands found")
	}

	for i := 0; i < len(cmds)-1; i++ {
		stdout, err := cmds[i].StdoutPipe()
		if err != nil {
			return "", fmt.Errorf("failed to create pipe: %w", err)
		}
		cmds[i+1].Stdin = stdout
	}

	var output bytes.Buffer
	cmds[len(cmds)-1].Stdout = &output

	var stderr bytes.Buffer
	cmds[len(cmds)-1].Stderr = &stderr

	for i := len(cmds) - 1; i > 0; i-- {
		if err := cmds[i].Start(); err != nil {
			return "", fmt.Errorf("failed to start command %d: %w", i, err)
		}
	}

	if err := cmds[0].Run(); err != nil {
		return "", fmt.Errorf("failed to run first command: %w: %s", err, stderr.String())
	}

	for i := 1; i < len(cmds); i++ {
		if err := cmds[i].Wait(); err != nil {
			return "", fmt.Errorf("failed to wait for command %d: %w: %s", i, err, stderr.String())
		}
	}

	return output.String(), nil
}

func GetShell() string {

	shell := os.Getenv("SHELL")
	shellName := strings.Split(shell, "/")[len(strings.Split(shell, "/"))-1]

	return shellName
}
