package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	//initialize sqlite
	db,err = initSQLite()
	if err != nil{
		return fmt.Errorf("error initializating sqlite: %v",err)
	}

	return nil
}

func GetSQLite() *gorm.DB{
	return db
}

func GetLogger(p string) *Logger {
	// initializer Logger
	logger = NewLogger(p)
	return logger
}
