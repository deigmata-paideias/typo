package root

import (
	"github.com/spf13/cobra"

	"github.com/deigmata-paideias/typo/internal/cmd"
)

func GetRootCommand() *cobra.Command {

	c := &cobra.Command{
		Use:   "Typo",
		Short: "Typo CLI Tool",
		Long:  "Like thefuck, but he uses Go to implement it more intelligently.",
	}

	c.AddCommand(cmd.VersionCommand())
	c.AddCommand(cmd.TypoCommand())
	c.AddCommand(cmd.ScanCommand())

	return c
}
