package utils

import (
	"os/exec"
	"strings"

	"github.com/deigmata-paideias/typo/internal/types"
)

// 检查命令是否存在于系统中
// c onefetch alias
// which c --> c: aliased to onefetch
// command -v c --> alias c=onefetch
func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func Convert(val, source string) ([]types.Command, error) {

	var commands []types.Command

	switch source {
	case "alias":
		lines := strings.Split(val, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || !strings.Contains(line, "=") {
				continue
			}

			parts := strings.SplitN(line, "=", 2)
			if len(parts) != 2 {
				continue
			}

			aliasName := strings.TrimSpace(parts[0])
			aliasValue := strings.Trim(strings.TrimSpace(parts[1]), `"' `)

			if aliasName == "" || aliasValue == "" {
				continue
			}

			command := types.Command{
				Name:        aliasName,
				Type:        string(types.Alias),
				Source:      "alias",
				Description: "Alias for: " + aliasValue,
			}
			commands = append(commands, command)
		}

	case "git":
		lines := strings.Split(val, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || !strings.Contains(line, "alias.") {
				continue
			}

			// 处理格式：alias.br=branch
			parts := strings.SplitN(line, "=", 2)
			if len(parts) != 2 {
				continue
			}

			aliasKey := strings.TrimSpace(parts[0])
			aliasValue := strings.Trim(strings.TrimSpace(parts[1]), `"' `)

			if aliasKey == "" || aliasValue == "" {
				continue
			}

			// 提取alias名称，去掉"alias."前缀
			if !strings.HasPrefix(aliasKey, "alias.") {
				continue
			}
			aliasName := strings.TrimPrefix(aliasKey, "alias.")

			command := types.Command{
				Name:        aliasName,
				Type:        string(types.Alias),
				Source:      "git",
				Description: "Git alias: " + aliasValue,
			}
			commands = append(commands, command)
		}

	case "man":
		lines := strings.Split(val, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			// 查找第一个 - 作为分隔符
			dashIndex := strings.Index(line, " - ")
			if dashIndex == -1 {
				continue
			}

			// 提取命令名部分
			cmdPart := strings.TrimSpace(line[:dashIndex])
			description := strings.TrimSpace(line[dashIndex+3:])

			// 分割多个命令名
			cmdNames := strings.Split(cmdPart, ", ")
			for _, rawCmd := range cmdNames {
				rawCmd = strings.TrimSpace(rawCmd)

				// 只处理 section 1 的命令
				if !strings.HasSuffix(rawCmd, "(1)") {
					continue
				}

				// 移除 (1) 后缀
				cmdName := strings.TrimSuffix(rawCmd, "(1)")
				if cmdName == "" {
					continue
				}

				command := types.Command{
					Name:        cmdName,
					Type:        string(types.Man),
					Source:      "man",
					Description: description,
				}

				// 检查命令是否实际存在
				if commandExists(cmdName) {
					commands = append(commands, command)
				}

				// 处理子命令，格式为：命令-子命令，git-branch, curl-config 等
				if strings.Contains(cmdName, "-") {
					parts := strings.SplitN(cmdName, "-", 2)
					if len(parts) == 2 {
						mainCmd := parts[0]
						subCmd := parts[1]
						if mainCmd != "" && subCmd != "" {
							// 标记为子命令，后续会处理
							commands = append(commands, types.Command{
								Name:        cmdName, // 完整名称
								Type:        string(types.Man),
								Source:      "man-subcommand",
								Description: description,
							})
						}
					}
				}
			}
		}
	}

	return commands, nil
}
