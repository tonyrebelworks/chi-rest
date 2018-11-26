package main

import (
	api "chi-rest/server/handler"

	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"
)

func registerRoutes(db *mgo.Session, r *chi.Mux) {
	handlerType := api.Handler{DB: db, Cfg: cfg}

	r.Route("/v1", func(r chi.Router) {

		member := api.MemberHandler{handlerType}

		r.Get("/me", member.Profile)
		r.Post("/register", member.RegisterHandler)
	})
}
