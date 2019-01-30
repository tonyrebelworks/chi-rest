package main

import (
	"bytes"
	apiHandler "chi-rest/server/handler"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	db *mgo.Session
)

type MyResponseWriter struct {
	http.ResponseWriter
	buf *bytes.Buffer
}

func registerMiddleware(dbSession *mgo.Session, r *chi.Mux) {
	db = dbSession
	r.Use(notfoundMiddleware)
	if debug {
		r.Use(middleware.Logger)
	}

	r.Use(middleware.Recoverer)
	r.Use(requestLoggerMiddleware)
}

func notfoundMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tctx := chi.NewRouteContext()
		rctx := chi.RouteContext(r.Context())

		if !rctx.Routes.Match(tctx, r.Method, r.URL.Path) {
			apiHandler.RespondWithJSON(w, 404, 404, "Request Not Found!", []map[string]interface{}{})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func requestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			err     error
			payload []byte
		)

		if r.Method == "GET" || r.Method == "HEAD" {
			next.ServeHTTP(w, r)
			return
		}

		if r.Method != "GET" {
			type requestLoggerEntity struct {
				id        bson.ObjectId `bson:"_id,omitempty"`
				Method    string        `bson:"method"`
				Payload   string        `bson:"payload"`
				CreatedAt time.Time     `bson:"created_at"`
			}

			payload, err = ioutil.ReadAll(r.Body)
			if err != nil {
				apiHandler.SendBadRequest(w, err.Error())
				return
			}

			rlE := requestLoggerEntity{
				Method:    r.Method,
				Payload:   string(payload),
				CreatedAt: time.Now(),
			}

			ds := db.Copy()
			defer ds.Close()
			err = ds.DB(cfg.GetString("database.mongo.db")).C("request_logger").Insert(&rlE)
			if err != nil {
				apiHandler.SendBadRequest(w, err.Error())
				return
			}

			r.Body = ioutil.NopCloser(bytes.NewBuffer(payload))

			// Create a response wrapper:
			mrw := &MyResponseWriter{
				ResponseWriter: w,
				buf:            &bytes.Buffer{},
			}

			// Call next handler, passing the response wrapper:
			next.ServeHTTP(mrw, r)
		}
	})
}
