package main

import (
	"chi-rest/lib"
	"chi-rest/model/mongo"
	"log"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/go-chi/chi"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
	cfg        lib.Config
	debug      = false
	host       string
)

func init() {
	cfg = lib.NewViperConfig(basepath)
	host = cfg.GetString("app.host")
	if cfg.GetBool("app.debug") {
		debug = true
		log.Printf("Running on Debug Mode: On at host [%v]", host)
	}
}

func main() {
	r := chi.NewRouter()

	// setup database connection
	dbInfo := mongo.Info{
		Host: cfg.GetString("database.mongo.host"),
		Db:   cfg.GetString("database.mongo.db"),
		User: cfg.GetString("database.mongo.user"),
		Pass: cfg.GetString("database.mongo.pass"),
	}
	db, err := dbInfo.Connect()
	if err != nil {
		panic(err)
	}

	// register middleware and application route
	registerMiddleware(r)
	registerRoutes(db, r)

	// run the app
	http.ListenAndServe(host, r)
}
