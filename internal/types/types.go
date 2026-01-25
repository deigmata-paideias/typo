package types

import "gorm.io/gorm"

type CommandType string

const (
	Alias CommandType = "alias"
	Man   CommandType = "man"
)

type ModeType string

const (
	Local ModeType = "local"
	LLM   ModeType = "llm"
)

type Config struct {
	Mode ModeType `yaml:"mode"`
	LLM  struct {
		Model   string `yaml:"model"`
		ApiKey  string `yaml:"api_key"`
		BaseUrl string `yaml:"base_url"`
	} `yaml:"llm"`
	Local struct {
		DBPath string `yaml:"db_path"`
	} `yaml:"local"`
}

// MatchResult 表示匹配结果
type MatchResult struct {
	Command string
	Desc    string
	Score   float64
}

// Command 命令
type Command struct {
	gorm.Model

	Name        string `gorm:"not null"`
	Type        string `gorm:"not null"` // system/alias/custom
	Source      string // man/alias/git/custom
	Description string
	Aliases     []CommandAlias
	Options     []CommandOption
}

// CommandAlias 别名
type CommandAlias struct {
	gorm.Model

	CommandID  uint   `gorm:"not null;index"`
	AliasName  string `gorm:"not null"`
	AliasValue string `gorm:"not null"`
	Source     string // alias/git/zsh/bash
	ShellType  string // bash/zsh/fish
	Priority   int    `gorm:"default:0"`
	Command    Command
}

// CommandOption 命令选项
type CommandOption struct {
	gorm.Model

	CommandID   uint `gorm:"not null;index"`
	OptionName  string
	Description string
	Command     Command
}

func (c *Command) TableName() string {
	return "commands"
}

func (ca *CommandAlias) TableName() string {
	return "command_aliases"
}

func (co *CommandOption) TableName() string {
	return "command_options"
}
