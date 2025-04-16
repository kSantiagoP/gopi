// originalmente pretendia usar mongodb, mas mudei de ideia por que já tinha importado gorm
// e por que bancos relacionais são melhores para dados estruturados e organizados como os que inseriremos
package config

import (
	"os"

	"github.com/kSantiagoP/gopi/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	//check existing db
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	//create db and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	//migrate schema
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	//return db only if everything is fine
	return db, nil
}
