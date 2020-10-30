## This is a simple implementation of Djangos Polls rest api written in golang.

# To run this.
    ` source utils/setenv.sh `
    ` go run cmd/migrations/main.go migrate <migrates all the models specified in database/migrations.go modelsToMigrate list>`
    ` go run cmd/migrations/main.go dropdb <deletes all the tables in the db so you can start fresh.>`
# TODO
    * Routing.
    * Handlers  (code logic making cruds.).
    * Database operations. (code logic in terms of database operations.).
    * Configuration (DB, prod/dev, ports).
    * Figure out Serialization.
