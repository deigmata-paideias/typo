package typo

import (
	"github.com/deigmata-paideias/typo/internal/repository"
	"github.com/deigmata-paideias/typo/internal/typo"
)

func RunTypo() error {

	repo := repository.NewRepository()
	typoInstance := typo.NewLocalTypo(repo)
	typoInstance.Typo("test")

	return nil
}
