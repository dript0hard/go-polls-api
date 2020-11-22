package main

import (
	"net/http"

	"github.com/dript0hard/pollsapi/config"
	"github.com/dript0hard/pollsapi/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	if config.DEBUG == true {
		r.Use(middleware.Logger)
	}

	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/", handlers.AuthRouter())
	r.Mount("/polls", handlers.PollRouter())

	r.Mount("/users", handlers.Test())

	http.ListenAndServe(":8080", r)
}
