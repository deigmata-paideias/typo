package custom

import (
	"github.com/deigmata-paideias/typo/internal/scanner"
	"github.com/deigmata-paideias/typo/internal/utils"
)

// GitAliasScanner 配置的 git 别名
type GitAliasScanner struct {
}

func NewGitAliasScanner() scanner.IScanner {

	return &GitAliasScanner{}
}

func (g *GitAliasScanner) Scan() (string, error) {

	// Github alias config, in ~/.gitconfig [alias]
	gitOutput, err := utils.ExecPipeCommand("git config --list | grep alias")
	if err != nil {
		return "", err
	}

	return gitOutput, nil
}
