package main

import (
	"fmt"

    "github.com/dript0hard/pollsapi/models"
    "github.com/dript0hard/pollsapi/database"
	"github.com/dript0hard/pollsapi/utils/password"
)

func main() {
    pwd := "my secret password"
    passwordHasher := password.NewPasswordSha512()
    hash := passwordHasher.HashPassword(pwd)
    fmt.Println(hash.String())
    passHash := hash.String()
    user := models.User{Username:"dn3i", Email:"deni1myftiu@gmail.com", Password: passHash}
    db, _ :=  database.OpenDB()
    db.Create(&user)
}
