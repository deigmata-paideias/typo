package scanner

import (
	"fmt"

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

	output, err := utils.ExecShell("alias")
	if err != nil {
		return "", err
	}

	fmt.Println("输出：", output)

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
