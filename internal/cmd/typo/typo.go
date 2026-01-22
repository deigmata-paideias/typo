package typo

import "github.com/deigmata-paideias/typo/internal/typo"

func RunTypo() error {

	typoInstance := typo.NewTypo()
	typoInstance.Typo("test")

	return nil
}
