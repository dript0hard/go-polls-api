package database

import (
	"gorm.io/gorm"
    "github.com/dript0hard/pollsapi/models"
)

func MigrateDB(db *gorm.DB) error {
    userErr := db.AutoMigrate(&models.User{})

    if userErr != nil {
        return userErr
    }

    pollErr := db.AutoMigrate(&models.Poll{})

    if pollErr != nil {
        return pollErr
    }

    choicesErr := db.AutoMigrate(&models.Choice{})

    if choicesErr != nil {
        return choicesErr
    }

    voteErr := db.AutoMigrate(&models.Vote{})

    if voteErr != nil {
        return voteErr
    }

    return nil
}
