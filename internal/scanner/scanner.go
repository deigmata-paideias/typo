package scanner

import (
	"github.com/deigmata-paideias/typo/internal/utils"
)

type IScanner interface {
	Scan() (string, error)
}

type AliasScanner struct {
}

func NewAliasScanner() IScanner {
	return &AliasScanner{}
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
}

func NewManScanner() IScanner {
	return &ManScanner{}
}

func (m *ManScanner) Scan() (string, error) {

	return "", nil
}
