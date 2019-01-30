package main

import (
	"github.com/globalsign/mgo"

	appMW "chi-rest/server/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	db *mgo.Session
)

func registerMiddleware(dbSession *mgo.Session, r *chi.Mux) {
	db = dbSession
	if debug {
		r.Use(middleware.Logger)
	}

	r.Use(middleware.Recoverer)

	// register app middleware
	mw := appMW.NewAppMiddleware(db, cfg)

	r.Use(mw.NotfoundMiddleware)
	r.Use(mw.RequestLoggerMiddleware)
}
