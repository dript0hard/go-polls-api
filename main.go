package main

import (
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/render"
    "github.com/dript0hard/pollsapi/handlers"
)

func main() {
    r := chi.NewRouter()
    r.Use(render.SetContentType(render.ContentTypeJSON))
    r.Mount("/", handlers.AuthRouter())
    http.ListenAndServe(":8080", r)
    // pwd := "password"
    // passwordHasher := password.NewPasswordSha512()
    // passHash := passwordHasher.HashPassword(pwd)
    // fmt.Println(passHash.String())
    // user := models.User{Username:"d3ni", Email:"deni1myftiu@gmail.com", Password: passHash.String()}
    // db, _ :=  database.OpenDB()
    // if result := db.Create(&user); result.Error != nil {
    //     fmt.Println(result.Error.Error())
    // }
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
