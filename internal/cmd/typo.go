package cmd

import (
	"github.com/spf13/cobra"

	"github.com/deigmata-paideias/typo/internal/cmd/typo"
)

func TypoCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "run",
		Aliases: []string{"run", "r"},
		Short:   "Each time it displays the 5 most probable commands; you can then use the up/down arrow keys to select and execute one.",
		RunE: func(_ *cobra.Command, _ []string) error {
			return typo.RunTypo()
		},
	}

	return cmd
}
