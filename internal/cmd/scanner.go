package cmd

import (
	"fmt"

	"github.com/deigmata-paideias/typo/internal/scanner"
	"github.com/deigmata-paideias/typo/internal/scanner/custom"
	"github.com/spf13/cobra"
)

type commandType string

const (
	alias commandType = "alias"
	man   commandType = "man"
)

func ScanCommand() *cobra.Command {

	var types string

	cmd := &cobra.Command{
		Use:     "scanner",
		Aliases: []string{"scanner", "s"},
		Short:   "Scanner command and save to db, container alias & man",
		RunE: func(_ *cobra.Command, _ []string) error {
			return runScanner(commandType(types))
		},
	}

	cmd.PersistentFlags().StringVarP(&types, "type", "t", "alias", "scanner command save to db, input alias or man.")

	return cmd
}

func runScanner(types commandType) error {

	switch types {
	case alias:
		aliasScanner := scanner.NewAliasScanner()
		output, err := aliasScanner.Scan()
		if err != nil {
			return err
		}
		fmt.Println(output)
		// add custom
		gitAliasScanner := custom.NewGitAliasScanner()
		gitOutput, err := gitAliasScanner.Scan()
		if err != nil {
			return err
		}
		fmt.Println(gitOutput)
	case man:
		manScanner := scanner.NewManScanner()
		output, err := manScanner.Scan()
		if err != nil {
			return err
		}
		fmt.Println(output)
	default:
		fmt.Println("not support")
	}

	return nil
}
