package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/deigmata-paideias/typo/internal/types"
)

// DefaultConfig returns the default configuration
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

// LoadConfig loads the configuration file
func LoadConfig() (*types.Config, error) {

	// Get default configuration first
	config := DefaultConfig()

	// Find configuration file path
	configPath, err := findConfigFile()
	if err != nil {
		// If configuration file not found, use default configuration
		return config, nil
	}

	// Read configuration file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("read config file failed: %w", err)
	}

	// Parse YAML
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("parse config file failed: %w", err)
	}

	// Expand ~ in path
	config.Local.DBPath = expandPath(config.Local.DBPath)

	return config, nil
}

// findConfigFile finds the configuration file
// Priority: current directory > ~/.config/typo/typo.config.yaml > /etc/typo/typo.config.yaml
func findConfigFile() (string, error) {
	homeDir, _ := os.UserHomeDir()

	// Possible configuration file paths
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
