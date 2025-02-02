package database

import (
	"daistant-core/configs"
	"daistant-core/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(config *configs.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.DB.SQLiteFilePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.ThirdPartyConnection{})
}
