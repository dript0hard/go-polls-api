package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbName     string = os.Getenv("DB_NAME")
	dbUserName string = os.Getenv("DB_USERNAME")
	dbPassword string = os.Getenv("DB_PASSWORD")
	dsn        string = fmt.Sprintf("user=%s password=%s dbname=%s",
		dbUserName, dbPassword, dbName)
)

func OpenDB() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
