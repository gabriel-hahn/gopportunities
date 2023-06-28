package config

import (
	"os"

	"github.com/gabriel-hahn/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createDbFile(dbPath string) error {
	err := os.MkdirAll("./db", os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(dbPath)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		err = createDbFile(dbPath)
		if err != nil {
			return nil, err
		}
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
