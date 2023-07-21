package config

import (
	"gorm.io/gorm"
)

var (
	logger *Logger
	db     *gorm.DB
)

func Init() error {
	var err error

	db, err = InitializeSQLite()
	if err != nil {
		logger.Errorf("error initializing sqlite: %v", err)
		return err
	}
	return nil
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}

func GetSQLite() *gorm.DB {
	return db
}
