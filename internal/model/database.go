package model

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB(storagePath string) error {
	const (
		op                 = "storage.sqlite.New"
		defaultStoragePath = "database.db"
	)

	var err error

	if len(storagePath) == 0 {
		storagePath = defaultStoragePath
	}

	db, err = gorm.Open(sqlite.Open(storagePath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	err = db.AutoMigrate(URL{})

	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	return nil
}
