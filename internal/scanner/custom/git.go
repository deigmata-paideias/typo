package custom

import (
	"github.com/deigmata-paideias/typo/internal/repository"
	"github.com/deigmata-paideias/typo/internal/scanner"
	"github.com/deigmata-paideias/typo/internal/utils"
)

// GitAliasScanner 配置的 git 别名
type GitAliasScanner struct {
	repo repository.IRepository
}

func NewGitAliasScanner(repo repository.IRepository) scanner.IScanner {

	return &GitAliasScanner{
		repo: repo,
	}
}

func (g *GitAliasScanner) Scan() (string, error) {

	// Github alias config, in ~/.gitconfig [alias]
	gitOutput, err := utils.ExecPipeCommand("git config --list | grep alias")
	if err != nil {
		return "", err
	}

	if gitOutput == "" {
		// 没有 git alias 配置
		return "", nil 
	}

	commands, err := utils.Convert(gitOutput, "git")
	if err != nil {
		return "", err
	}

	// 保存到数据库
	if err := g.repo.BatchInsertCommand(commands); err != nil {
		return "", err
	}

	return "", nil
}
