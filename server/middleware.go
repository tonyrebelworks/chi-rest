package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func registerMiddleware(r *chi.Mux) {
	if debug {
		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
	}

	r.Use(middleware.Recoverer)
}
