package scanner

import (
	"github.com/deigmata-paideias/typo/internal/repository"
	"github.com/deigmata-paideias/typo/internal/utils"
)

type IScanner interface {
	Scan() (string, error)
}

type AliasScanner struct {
	repository.IRepository
}

func NewAliasScanner(repo repository.IRepository) IScanner {

	return &AliasScanner{
		repo,
	}
}

func (a *AliasScanner) Scan() (string, error) {

	// System alias, contain zsh and bash etc.
	output, err := utils.ExecCommand("zsh", "-i", "-c", "alias")
	if err != nil {
		return "", err
	}

	return output, nil
}

type ManScanner struct {
	repository.IRepository
}

func NewManScanner(repo repository.IRepository) IScanner {

	return &ManScanner{
		repo,
	}
}

func (m *ManScanner) Scan() (string, error) {

	return "", nil
}
