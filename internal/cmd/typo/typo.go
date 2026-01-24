package typo

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/deigmata-paideias/typo/internal/repository"
	"github.com/deigmata-paideias/typo/internal/scanner/history"
	"github.com/deigmata-paideias/typo/internal/tui"
	"github.com/deigmata-paideias/typo/internal/typo"
)

func RunTypo() error {

	repo := repository.NewRepository()
	historyScanner := history.NewHistory()
	typoInstance := typo.NewLocalTypo(repo, historyScanner)

	originalCmd, matches, err := typoInstance.Typo()
	if err != nil {
		return err
	}

	if len(matches) == 0 {
		fmt.Printf("命令正确: %s\n", originalCmd)
		return nil
	}

	// 用 TUI 界面选
	selectedCmd, err := tui.RunSelector(originalCmd, matches)
	if err != nil {
		return err
	}

	if selectedCmd != "" {
		parts := strings.Split(originalCmd, " ")
		if len(parts) > 0 {
			parts[0] = selectedCmd
			correctedCmd := strings.Join(parts, " ")

			// 执行修正后的命令
			cmd := exec.Command("zsh", "-c", correctedCmd)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				var e *exec.ExitError
				if errors.As(err, &e) {
					// 原样返回退出码
					os.Exit(e.ExitCode())
				}
				// 启动失败、找不到命令或者其他错误
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	}

	return nil
}
