package config

import (
	"os"

	"github.com/levifleal/MyFlix/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initSQLite() (*gorm.DB, error) {
	logger := GetLogger("SQLite")
	dbPath := "./db/main.db"

	//Check if the DB file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	//Create DB and Connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	//Migrate the Schema
	err = db.AutoMigrate(&schemas.Content{})
	if err != nil {
		logger.Errorf("sqlite autoMigration error: %v", err)
	}

	//Return the DB
	return db,nil
}
