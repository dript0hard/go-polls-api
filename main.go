package main

import (
	"fmt"

    "github.com/dript0hard/pollsapi/models"
    "github.com/dript0hard/pollsapi/database"
	"github.com/dript0hard/pollsapi/utils/password"
)

func main() {
    pwd := "password"
    passwordHasher := password.NewPasswordSha512()
    passHash := passwordHasher.HashPassword(pwd)
    fmt.Println(passHash.String())
    user := models.User{Username:"d3ni", Email:"deni1myftiu@gmail.com", Password: passHash.String()}
    db, _ :=  database.OpenDB()
    if result := db.Create(&user); result.Error != nil {
        fmt.Println(result.Error.Error())
    }
    // ======================================
    // db, _ :=  database.OpenDB()
    // user := models.User{}
    // db.First(&user)
    // hr := password.NewHashResult(user.Password)
    // fmt.Printf("%#v\n ", hr)

    // passwordHasher := password.NewPasswordSha512()
    // valid := passwordHasher.ValidatePassword(pwd, hr)
    // fmt.Printf("%#v\n ", valid)
}
