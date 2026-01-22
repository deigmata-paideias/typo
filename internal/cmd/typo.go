package cmd

import (
	"github.com/spf13/cobra"

	"github.com/deigmata-paideias/typo/internal/cmd/typo"
)

func TypoCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "run",
		Aliases: []string{"run", "r"},
		Short:   "Run typo to fix your command.",
		RunE: func(_ *cobra.Command, _ []string) error {
			return typo.RunTypo()
		},
	}

	return cmd
}
