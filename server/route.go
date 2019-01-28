package main

import (
	api "chi-rest/server/handler"

	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"
)

func registerRoutes(db *mgo.Session, r *chi.Mux) {
	handlerType := api.Handler{DB: db, Cfg: cfg}

	r.Route("/v1", func(r chi.Router) {

		m := api.MemberHandler{handlerType}

		r.Get("/me", m.Profile)
		r.Post("/register", m.RegisterHandler)
	})
}
