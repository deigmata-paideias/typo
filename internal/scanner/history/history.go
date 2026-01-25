package history

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/deigmata-paideias/typo/internal/scanner"
	"github.com/deigmata-paideias/typo/internal/utils"
)

type History struct {
}

func NewHistory() scanner.IScanner {
	return &History{}
}

func (h *History) Scan() (string, error) {

	homeDir := os.Getenv("HOME")

	switch utils.GetShell() {
	case "bash":
		return bash(homeDir)
	case "zsh":
		return zsh(homeDir)
	default:
		return "", fmt.Errorf("unsupported shell: %s", utils.GetShell())
	}
}

func zsh(home string) (string, error) {

	zshHistory := filepath.Join(home, ".zsh_history")

	if file, err := os.Open(zshHistory); err == nil {
		defer file.Close()

		// Read the last two lines, the previous line in zsh is the current
		// command execution line, the next line is the previous command ?
		var lastCmd string
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			lastCmd = scanner.Text()
		}

		// zsh history format: ": timestamp:0;command"
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
