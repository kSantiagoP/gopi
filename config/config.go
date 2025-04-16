package config

//Lógica para configurações gerais e do banco de dados

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

	db, err = InitializeSQlite()

	if err != nil {
		return fmt.Errorf("error initializing database: %v", err)
	}

	return nil
}

func GetSQlite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	//initialize logger
	logger = NewLogger(p)
	return logger
}
