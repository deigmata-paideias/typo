package scanner

import (
	"strings"

	"github.com/deigmata-paideias/typo/internal/repository"
	"github.com/deigmata-paideias/typo/internal/types"
	"github.com/deigmata-paideias/typo/internal/utils"
)

type IScanner interface {
	Scan() (string, error)
}

// alias command

type AliasScanner struct {
	repo        repository.IRepository
	CustomAlias []IScanner
}

func NewAliasScanner(repo repository.IRepository, sc []IScanner) IScanner {

	return &AliasScanner{
		repo:        repo,
		CustomAlias: sc,
	}
}

func (a *AliasScanner) Scan() (string, error) {

	// System alias, contain zsh and bash etc.
	output, err := utils.ExecCommand("zsh", "-i", "-c", "alias")
	if err != nil {
		return "", err
	}

	commands, err := utils.Convert(output, "alias")
	if err != nil {
		return "", err
	}

	// Save to database
	if err := a.repo.BatchInsertCommand(commands); err != nil {
		return "", err
	}

	for _, scFunc := range a.CustomAlias {
		_, err = scFunc.Scan()
		if err != nil {
			return "", err
		}
	}

	return "", nil
}

// man command

type ManScanner struct {
	repo       repository.IRepository
	manCommand []types.Command
}

func NewManScanner(repo repository.IRepository) IScanner {

	return &ManScanner{
		repo:       repo,
		manCommand: make([]types.Command, 0),
	}
}

func (m *ManScanner) Scan() (string, error) {

	output, err := utils.ExecCommand("man", "-k", ".")
	if err != nil {
		return "", err
	}

	commands, err := utils.Convert(output, "man")
	if err != nil {
		return "", err
	}

	// Separate normal commands and subcommands
	var normalCommands []types.Command
	var subCommands []types.Command

	for _, cmd := range commands {
		if cmd.Source == "man-subcommand" {
			subCommands = append(subCommands, cmd)
		} else {
			normalCommands = append(normalCommands, cmd)
		}
	}

	// Save normal commands to database
	if err := m.repo.BatchInsertCommand(normalCommands); err != nil {
		return "", err
	}

	// Process subcommands
	if err := m.processSubCommands(subCommands); err != nil {
		return "", err
	}

	return "", nil
}

// processSubCommands handles all subcommands generically
func (m *ManScanner) processSubCommands(commands []types.Command) error {
	// Group subcommands by main command
	subCommandsByMain := make(map[string][]types.CommandOption)

	// Collect all subcommands
	for _, cmd := range commands {
		if cmd.Source == "man-subcommand" && strings.Contains(cmd.Name, "-") {
			parts := strings.SplitN(cmd.Name, "-", 2)
			if len(parts) == 2 {
				mainCmd := parts[0]
				subCmd := parts[1]
				if mainCmd != "" && subCmd != "" {
					subCommandsByMain[mainCmd] = append(subCommandsByMain[mainCmd], types.CommandOption{
						OptionName:  subCmd,
						Description: cmd.Description,
					})
				}
			}
		}
	}

	// Process subcommands for each main command
	for mainCmd, subCommands := range subCommandsByMain {
		// Find the main command
		mainCommand, err := m.repo.FindCommandByName(mainCmd)
		if err != nil {
			// If the main command doesn't exist, create one
			newMainCommand := types.Command{
				Name:        mainCmd,
				Type:        string(types.Man),
				Source:      "man",
				Description: "Command with subcommands",
			}
			if err := m.repo.BatchInsertCommand([]types.Command{newMainCommand}); err != nil {
				continue // Skip errors, continue processing other commands
			}
			mainCommand, _ = m.repo.FindCommandByName(mainCmd)
		}

		// Set CommandID for each subcommand and insert
		var optionsWithID []types.CommandOption
		for _, subCmd := range subCommands {
			subCmd.CommandID = mainCommand.ID
			optionsWithID = append(optionsWithID, subCmd)
		}

		if len(optionsWithID) > 0 {
			err = m.repo.BatchInsertCommandOptions(optionsWithID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
