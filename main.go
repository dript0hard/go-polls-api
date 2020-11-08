package main

import (
	"net/http"

	"github.com/dript0hard/pollsapi/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Logger)

	r.Mount("/", handlers.AuthRouter())
	r.Mount("/users", handlers.Test())
	r.Mount("/polls", handlers.PollRouter())

	http.ListenAndServe(":8080", r)
}
