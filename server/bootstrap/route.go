package bootstrap

import (
	"chi-rest/server/handler"
	"chi-rest/usecase"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// RegisterRoutes all routes for the apps
func (app App) RegisterRoutes() {
	h := handler.Contract{
		DB: app.DB,
	}
	app.R.Route("/v1", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, req *http.Request) {
			fmt.Println("This is message from get function [/] => json config: " + app.Config.GetString("app.host"))

			err := usecase.UC{DB: h.DB}.GetData()
			if err != nil {
				fmt.Println(err)
			}
			h.SendSuccess(w, map[string]interface{}{}, nil)
			return
		})
	})
}
