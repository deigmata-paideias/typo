package cmd

import (
	"github.com/deigmata-paideias/typo/internal/typo"
	"github.com/spf13/cobra"
)

func TypoCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "run",
		Aliases: []string{"run", "r"},
		Short:   "Run typo to fix your command.",
		RunE: func(_ *cobra.Command, _ []string) error {
			return runTypo()
		},
	}

	return cmd
}

func runTypo() error {

	typoInstance := typo.NewTypo()
	typoInstance.Typo("test")

	return nil
}
