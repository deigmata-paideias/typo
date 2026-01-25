package custom

import (
	"github.com/deigmata-paideias/typo/internal/repository"
	"github.com/deigmata-paideias/typo/internal/scanner"
	"github.com/deigmata-paideias/typo/internal/utils"
)

// GitAliasScanner scans configured git aliases
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
		// No git alias configuration found
		return "", nil
	}

	commands, err := utils.Convert(gitOutput, "git")
	if err != nil {
		return "", err
	}

	// Save to database
	if err := g.repo.BatchInsertCommand(commands); err != nil {
		return "", err
	}

	return "", nil
}
