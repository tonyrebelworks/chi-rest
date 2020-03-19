package main

import (
	"chi-rest/lib/utils"
	"chi-rest/server/bootstrap"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/valve"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
	config     utils.Config
	debug      = false
	host       string
)

// EnvConfigPath environtment variable that set the config path
const EnvConfigPath = "REBEL_CLI_CONFIG_PATH"

func init() {
	configFile := os.Getenv(EnvConfigPath)
	if configFile == "" {
		configFile = "../config.json"
	}

	log.Println(configFile)

	config = utils.NewViperConfig(basepath, configFile)
	host = config.GetString("app.host")

	debug = config.GetBool("app.debug")
	if config.GetBool("app.debug") {
		debug = true
		log.Printf("Running on Debug Mode: On at host [%v]", host)
	}
}

func main() {
	// gracefull shutdown handler
	valv := valve.New()
	baseCtx := valv.Context()

	// start new app
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
			"X-SIGNATURE",
			"X-TIMESTAMPT",
		},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)
	if debug {
		r.Use(middleware.Logger)
	}
	r.Use(middleware.Recoverer)

	db := bootstrap.SetupDB(config)
	validator := bootstrap.SetupValidator(config)

	// loan the app skeleton
	app := bootstrap.App{
		R:         r,
		Config:    config,
		DB:        db,
		Validator: validator,
	}

	app.RegisterRoutes()

	// handle gracefull shutdown
	srv := http.Server{Addr: host, Handler: chi.ServerBaseContext(baseCtx, r)}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println("shutting down..")
			valv.Shutdown(20 * time.Second)
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()
			srv.Shutdown(ctx)
			select {
			case <-time.After(21 * time.Second):
				fmt.Println("not all connections done")
			case <-ctx.Done():

			}
		}
	}()
	srv.ListenAndServe()
}
