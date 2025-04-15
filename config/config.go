package config

//Lógica para configurações gerais e do banco de dados

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func Init() error {
	return nil
}
