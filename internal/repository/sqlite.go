package repository

import (
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/deigmata-paideias/typo/internal/types"
)

type IRepository interface {
	BatchInsertCommand(records []types.Command) error
	BatchInsertCommandAlias(aliases []types.CommandAlias) error
	BatchInsertCommandOptions(categories []types.CommandOption) error
	GetAllCommands() ([]types.Command, error)
	FindCommandByName(name string) (*types.Command, error)
	GetAllCommandNames() ([]string, error)
	GetCommandOptions(commandName string) ([]types.CommandOption, error)
	GetAllCommandOptionNames(commandName string) ([]string, error)
}

type Repository struct {
	db *gorm.DB
}

// NewRepository 创建 Sqlite 数据库连接，使用默认路径
func NewRepository() IRepository {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	dbPath := filepath.Join(homeDir, ".config", "typo", "typo.db")
	return NewRepositoryWithPath(dbPath)
}

// NewRepositoryWithPath 创建 Sqlite 数据库连接
func NewRepositoryWithPath(dbPath string) IRepository {

	dbDir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		panic(err)
	}

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}
	db, err := newSqliteDB(dbPath, config)
	if err != nil {
		panic(err)
	}

	return &Repository{
		db: db,
	}
}

func newSqliteDB(dsn string, config *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), config)
	if err != nil {
		return nil, err
	}

	// 自动迁移表结构
	if err := db.AutoMigrate(&types.CommandAlias{}); err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&types.CommandOption{}); err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&types.Command{}); err != nil {
		return nil, err
	}

	return db, nil
}

func (r Repository) BatchInsertCommand(records []types.Command) error {

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, record := range records {
		var count int64
		// 用 name 和 type 作为唯一键检查，可能存在多个 alias 指向同一个 command，但是 type 不同
		// alias c --> onefetch
		// git c --> git commit -m xxx
		if err := tx.Model(&types.Command{}).Where("name = ? and type = ?", record.Name, record.Type).Count(&count).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 如果不存在则插入
		if count == 0 {
			if err := tx.Create(&record).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

func (r Repository) BatchInsertCommandAlias(aliases []types.CommandAlias) error {

	return r.db.Create(aliases).Error
}

func (r Repository) BatchInsertCommandOptions(categories []types.CommandOption) error {

	return r.db.Create(categories).Error
}

func (r Repository) GetAllCommands() ([]types.Command, error) {

	var commands []types.Command

	err := r.db.Find(&commands).Error

	return commands, err
}

func (r Repository) FindCommandByName(name string) (*types.Command, error) {

	var command types.Command

	err := r.db.Where("name = ?", name).First(&command).Error
	if err != nil {
		return nil, err
	}

	return &command, nil
}

func (r Repository) GetAllCommandNames() ([]string, error) {

	var names []string

	err := r.db.Model(&types.Command{}).Pluck("name", &names).Error

	return names, err
}

// GetCommandOptions 获取指定命令的所有选项/子命令
func (r Repository) GetCommandOptions(commandName string) ([]types.CommandOption, error) {
	var (
		command types.Command
		options []types.CommandOption
	)

	// 先找到对应的命令
	err := r.db.Where("name = ?", commandName).First(&command).Error
	if err != nil {
		return nil, err
	}

	// 获取该命令的所有选项
	err = r.db.Where("command_id = ?", command.ID).Find(&options).Error
	return options, err
}

// GetAllCommandOptionNames 获取指定命令的所有选项/子命令名称
func (r Repository) GetAllCommandOptionNames(commandName string) ([]string, error) {

	var (
		command types.Command
		names   []string
	)

	// 先找到对应的命令
	err := r.db.Where("name = ?", commandName).First(&command).Error
	if err != nil {
		return nil, err
	}

	// 获取该命令的所有选项名称
	err = r.db.Model(&types.CommandOption{}).
		Where("command_id = ?", command.ID).
		Pluck("option_name", &names).Error
	return names, err
}
