package main

import (
	"net/http"

	"github.com/dript0hard/pollsapi/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
    "github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

    r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
    r.Use(middleware.Logger)

	r.Mount("/", handlers.AuthRouter())
	r.Mount("/users", handlers.Test())
	http.ListenAndServe(":8080", r)
}
