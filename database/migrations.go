package database

import (
	"github.com/dript0hard/pollsapi/models"
	"gorm.io/gorm"
)

var (
	modelsToMigrate []interface{} = []interface{}{
		models.User{},
		models.Poll{},
		models.Choice{},
		models.Vote{},
	}
)

func MigrateDB(db *gorm.DB) error {
	for _, model := range modelsToMigrate {
		err := db.AutoMigrate(model)

		if err != nil {
			return err
		}
	}
	return nil
}

func DropDB(db *gorm.DB) error {
	for _, model := range modelsToMigrate {
		err := db.Migrator().DropTable(model)

		if err != nil {
			return err
		}
	}
	return nil
}
