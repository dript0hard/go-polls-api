package handlers

import (
	"github.com/go-chi/chi"
)

func AuthRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", loginHandler)
	r.Post("/register", registerHandler)

	return r
}

func PollRouter() chi.Router {
	r := chi.NewRouter()
	r.Post("/", createPoll)
	r.Get("/", listPolls)
	r.Route("/{pollId}", func(r chi.Router) {
		r.Get("/", getPollById)
		r.Put("/", updatePoll)
		r.Delete("/", deletePoll)
		r.Route("/choices", func(r chi.Router) {
			r.Get("/", listChoices)
			r.Post("/{choiceId}/vote", voteChoice)
		})
	})
	return r
}

