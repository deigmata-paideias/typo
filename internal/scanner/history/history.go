package history

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/deigmata-paideias/typo/internal/scanner"
)

type History struct {
}

func NewHistory() scanner.IScanner {
	return &History{}
}

func (h *History) Scan() (string, error) {

	shell := os.Getenv("SHELL")
	homeDir := os.Getenv("HOME")
	shellName := strings.Split(shell, "/")[len(strings.Split(shell, "/"))-1]

	switch shellName {
	case "bash":
		return bash(homeDir)
	case "zsh":
		return zsh(homeDir)
	default:
		return "", fmt.Errorf("unsupported shell: %s", shell)
	}
}

func zsh(home string) (string, error) {

	zshHistory := filepath.Join(home, ".zsh_history")

	if file, err := os.Open(zshHistory); err == nil {
		defer file.Close()

		// 读取最后两行，zsh 的前一行是当前
		// 命令运行的一行，后一行是上一条命令 ?
		var lastCmd string
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			lastCmd = scanner.Text()
		}

		// zsh history 格式: ": timestamp:0;command"
		if strings.Contains(lastCmd, ";") {
			//
			parts := strings.SplitN(lastCmd, ";", 2)
			return parts[len(parts)-1], nil
		}
		return lastCmd, nil
	}

	return "", fmt.Errorf("failed to open zsh history file: %s", zshHistory)
}

func bash(home string) (string, error) {

	bashHistory := filepath.Join(home, ".bash_history")
	if file, err := os.Open(bashHistory); err == nil {
		defer file.Close()

		var lastCmd string
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			lastCmd = scanner.Text()
		}
		return lastCmd, nil
	}

	return "", fmt.Errorf("failed to open bash history file: %s", bashHistory)
}
