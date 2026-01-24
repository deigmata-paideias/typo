package scanner

import (
	"github.com/deigmata-paideias/typo/internal/repository"
	"github.com/deigmata-paideias/typo/internal/types"
	"github.com/deigmata-paideias/typo/internal/utils"
)

type IScanner interface {
	Scan() (string, error)
}

// alias command

type AliasScanner struct {
	repo        repository.IRepository
	CustomAlias []IScanner
}

func NewAliasScanner(repo repository.IRepository, sc []IScanner) IScanner {

	return &AliasScanner{
		repo:        repo,
		CustomAlias: sc,
	}
}

func (a *AliasScanner) Scan() (string, error) {

	// System alias, contain zsh and bash etc.
	output, err := utils.ExecCommand("zsh", "-i", "-c", "alias")
	if err != nil {
		return "", err
	}

	commands, err := utils.Convert(output, "alias")
	if err != nil {
		return "", err
	}

	// 保存到数据库
	if err := a.repo.BatchInsertCommand(commands); err != nil {
		return "", err
	}

	for _, scFunc := range a.CustomAlias {
		_, err = scFunc.Scan()
		if err != nil {
			return "", err
		}
	}

	return "", nil
}

// man command

type ManScanner struct {
	repo       repository.IRepository
	manCommand []types.Command
}

func NewManScanner(repo repository.IRepository) IScanner {

	return &ManScanner{
		repo:       repo,
		manCommand: make([]types.Command, 0),
	}
}

func (m *ManScanner) Scan() (string, error) {

	output, err := utils.ExecCommand("man", "-k", ".")
	if err != nil {
		return "", err
	}

	commands, err := utils.Convert(output, "man")
	if err != nil {
		return "", err
	}

	// 保存到数据库
	if err := m.repo.BatchInsertCommand(commands); err != nil {
		return "", err
	}

	return "", nil
}
