package cmd

import (
	"github.com/deigmata-paideias/typo/internal/types"
	"github.com/spf13/cobra"

	"github.com/deigmata-paideias/typo/internal/cmd/scanner"
)

func ScanCommand() *cobra.Command {

	var t string

	cmd := &cobra.Command{
		Use:     "scanner",
		Aliases: []string{"scanner", "s"},
		Short:   "Scanner command and save to db, container alias & man",
		RunE: func(_ *cobra.Command, _ []string) error {
			return scanner.RunScanner(types.CommandType(t))
		},
	}

	cmd.PersistentFlags().StringVarP(&t, "type", "t", "alias", "scanner command save to db, input alias or man.")

	return cmd
}
