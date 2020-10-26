package main

import (
	"fmt"

	"github.com/dript0hard/pollsapi/database"
)

func main() {
    db, connErr := database.OpenDB()

    if connErr != nil {
        fmt.Println("Could not connect to the database.")
        fmt.Println(connErr.Error())
    }

    migrationErr := database.MigrateDB(db)

    if migrationErr != nil {
        fmt.Println("Could not migrate the database.")
        fmt.Println(migrationErr.Error())
    }
}
