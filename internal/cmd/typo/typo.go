package typo

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

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

	// 如果没有匹配项，说明命令可能是正确的
	if len(matches) == 0 {
		fmt.Printf("命令正确: %s\n", originalCmd)
		return nil
	}

	// 用 TUI 界面选择
	selectedCmd, err := tui.RunSelector(originalCmd, matches)
	if err != nil {
		return err
	}

	if selectedCmd != "" {
		// selectedCmd 已经是完整的修正命令，直接执行
		fmt.Println(selectedCmd)
		// 执行修正后的命令
		cmd := exec.Command("zsh", "-c", selectedCmd)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			var e *exec.ExitError
			if errors.As(err, &e) {
				// 原样返回退出码
				os.Exit(e.ExitCode())
			}
			// 启动失败、找不到命令等
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	return nil
}
