package main

import (
    "os"
	"fmt"

	"github.com/dript0hard/pollsapi/database"
)

var USAGE string = "usage:\n\tmigrations migrate - Migrates all tables.\n\tmigrations dropdb - Deletes all the teables in the db."

func main() {
    if len(os.Args) < 2 {
        fmt.Println(USAGE)
        return
    }

    db, connErr := database.OpenDB()

    if connErr != nil {
        fmt.Println("Could not connect to the database.")
        fmt.Println(connErr.Error())
    }

    switch os.Args[1] {
        case "migrate": {
            migrationErr := database.MigrateDB(db)

            if migrationErr != nil {
                fmt.Println("Could not migrate the database.")
                fmt.Println(migrationErr.Error())
            }
        }
        case "dropdb": {
            deleteErr := database.DropDB(db)

            if deleteErr != nil {
                fmt.Println("Could not delete all the tables in the database.")
                fmt.Println(deleteErr.Error())
            }
        }
        default: {
            fmt.Println(USAGE)
            return
        }
    }
}
