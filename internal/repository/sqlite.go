package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/deigmata-paideias/typo/internal/types"
)

type IRepository interface {
	BatchInsertCommandRecords(records []types.CommandRecord) error
	BatchInsertCommandAliases(aliases []types.CommandAlias) error
	BatchInsertCommandCategories(categories []types.CommandCategory) error
}

type Repository struct {
	db *gorm.DB
}

// NewRepository 创建 Sqlite 数据库连接
func NewRepository() IRepository {

	config := &gorm.Config{
		// 禁用 创建数据库外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	dsn := "file" + ":~/.config/typo/typo.db"

	db, err := newSqliteDB(dsn, config)
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

	err = db.AutoMigrate(types.CommandCategory{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(types.CommandRecord{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(types.CommandAlias{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r Repository) BatchInsertCommandRecords(records []types.CommandRecord) error {

	return r.db.Create(records).Error
}

func (r Repository) BatchInsertCommandAliases(aliases []types.CommandAlias) error {

	return r.db.Create(aliases).Error
}

func (r Repository) BatchInsertCommandCategories(categories []types.CommandCategory) error {

	return r.db.Create(categories).Error
}
