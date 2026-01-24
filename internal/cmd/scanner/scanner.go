package scanner

import (
	"fmt"

	"github.com/deigmata-paideias/typo/internal/repository"
	"github.com/deigmata-paideias/typo/internal/scanner"
	"github.com/deigmata-paideias/typo/internal/scanner/custom"
	"github.com/deigmata-paideias/typo/internal/types"
)

func RunScanner(t types.CommandType) error {

	switch t {
	case types.Alias:
		if err := execAliasScanner(); err != nil {
			return err
		}
	case types.Man:
		if err := execManScanner(); err != nil {
			return err
		}
	default:
		fmt.Println("not support")
	}

	return nil
}

func execAliasScanner() error {

	repo := repository.NewRepository()
	// custom plugin
	gitAliasScanner := custom.NewGitAliasScanner(repo)

	aliasScanner := scanner.NewAliasScanner(
		repo,
		[]scanner.IScanner{gitAliasScanner},
	)

	if _, err := aliasScanner.Scan(); err != nil {
		return err
	}

	return nil
}

func execManScanner() error {

	repo := repository.NewRepository()
	manScanner := scanner.NewManScanner(repo)

	if _, err := manScanner.Scan(); err != nil {
		return err
	}

	return nil
}
