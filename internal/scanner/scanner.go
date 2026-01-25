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

	// 保存到数据库
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

	// 分离普通命令和子命令
	var normalCommands []types.Command
	var subCommands []types.Command

	for _, cmd := range commands {
		if cmd.Source == "man-subcommand" {
			subCommands = append(subCommands, cmd)
		} else {
			normalCommands = append(normalCommands, cmd)
		}
	}

	// 保存普通命令到数据库
	if err := m.repo.BatchInsertCommand(normalCommands); err != nil {
		return "", err
	}

	// 处理子命令
	if err := m.processSubCommands(subCommands); err != nil {
		return "", err
	}

	return "", nil
}

// processSubCommands 通用处理所有子命令
func (m *ManScanner) processSubCommands(commands []types.Command) error {
	// 按主命令分组处理子命令
	subCommandsByMain := make(map[string][]types.CommandOption)

	// 收集所有子命令
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

	// 处理每个主命令的子命令
	for mainCmd, subCommands := range subCommandsByMain {
		// 查找主命令
		mainCommand, err := m.repo.FindCommandByName(mainCmd)
		if err != nil {
			// 如果主命令不存在，创建一个
			newMainCommand := types.Command{
				Name:        mainCmd,
				Type:        string(types.Man),
				Source:      "man",
				Description: "Command with subcommands",
			}
			if err := m.repo.BatchInsertCommand([]types.Command{newMainCommand}); err != nil {
				continue // 跳过错误，继续处理其他命令
			}
			mainCommand, _ = m.repo.FindCommandByName(mainCmd)
		}

		// 为每个子命令设置CommandID并插入
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
