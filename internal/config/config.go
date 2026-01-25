package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/deigmata-paideias/typo/internal/types"
)

// DefaultConfig 返回默认配置
func DefaultConfig() *types.Config {

	homeDir, _ := os.UserHomeDir()

  // ~/.config/typo/typo.db
	defaultDBPath := filepath.Join(homeDir, ".config", "typo", "typo.db")

	return &types.Config{
		Mode: types.Local,
		Local: struct {
			DBPath string `yaml:"db_path"`
		}{
			DBPath: defaultDBPath,
		},
		LLM: struct {
			Model   string `yaml:"model"`
			ApiKey  string `yaml:"api_key"`
			BaseUrl string `yaml:"base_url"`
		}{
			Model:   "gpt-3.5-turbo",
			ApiKey:  "",
			BaseUrl: "https://api.openai.com/v1",
		},
	}
}

// LoadConfig 加载配置文件
func LoadConfig() (*types.Config, error) {

	// 先获取默认配置
	config := DefaultConfig()

	// 查找配置文件路径
	configPath, err := findConfigFile()
	if err != nil {
		// 如果找不到配置文件，使用默认配置
		return config, nil
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("read config file failed: %w", err)
	}

	// 解析 YAML
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("parse config file failed: %w", err)
	}

	// 展开路径中的 ~
	config.Local.DBPath = expandPath(config.Local.DBPath)

	return config, nil
}

// findConfigFile 查找配置文件
// 优先级：当前目录 > ~/.config/typo/typo.config.yaml > /etc/typo/typo.config.yaml
func findConfigFile() (string, error) {
	homeDir, _ := os.UserHomeDir()

	// 可能的配置文件路径
	paths := []string{
		"typo.config.yaml",
		".typo.config.yaml",
		filepath.Join(homeDir, ".config", "typo", "typo.config.yaml"),
		"/etc/typo/typo.config.yaml",
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("config file not found")
}

func expandPath(path string) string {

  if strings.HasPrefix(path, "~/") {
		homeDir, _ := os.UserHomeDir()
		return filepath.Join(homeDir, path[2:])
	}

  return path
}
