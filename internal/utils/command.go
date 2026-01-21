package utils

import (
	"fmt"
	"os/exec"
)

// ExecShell 执行 shell 命令
func ExecShell(cmd string, args ...string) (string, error) {

	output, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil && len(output) > 0 {

		return "", fmt.Errorf("%w: %s", err, string(output))
	}

	return string(output), nil
}
