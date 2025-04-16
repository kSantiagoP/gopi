package config

//Lógica para configurações gerais e do banco de dados

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	//initialize logger
	logger = NewLogger(p)
	return logger
}
