package journeyplan

import (
	"chi-rest/bootstrap"
	"chi-rest/services/journeyplan/handler"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

// RegisterRoutes all routes for the apps
func RegisterRoutes(r *chi.Mux, app *bootstrap.App) {

	//The url pointing to API definition"
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(app.Config.GetString("app.app_host")+"/swagger/doc.json"),
	))

	h := handler.Contract{app}
	r.Route("/v1", func(r chi.Router) {
		// r.Get("/", h.Hello)

		//Journey CMS
		r.Get("/journey", h.GetAllJourney)
		r.Get("/journey/{code}", h.GetDetailJourney)
		r.Post("/journey", h.AddJourney)
		r.Put("/journey/{code}", h.UpdateJourney)
		r.Delete("/journey/{code}", h.DeleteJourney)
		r.Get("/journey/report/{code}", h.GetReportJourney)

		//Journey Mobile
		r.Get("/journeymobile", h.GetAllJourneyMobile)
		r.Get("/journeymobile/{code}", h.GetDetailJourneyMobile)
		r.Put("/journey/time", h.UpdateTimeJourney)
		r.Get("/journey/interval", h.GetInterval)
		r.Put("/journey/interval", h.UpdateInterval)
		r.Post("/journey/trackingtime", h.AddTrackingTimeJourney)
		r.Post("/journey/url", h.AddURLFirebase)

		//Competitor Analysis
		r.Get("/company", h.GetAllCompany)
		r.Get("/company/{code}", h.GetDetailCompany)
		r.Post("/company", h.AddCompany)
		r.Put("/company/{code}", h.UpdateCompany)
		r.Delete("/company/{code}", h.DeleteCompany)

		//Product
		r.Get("/product", h.GetAllProduct)
		r.Get("/product/{code}", h.GetDetailProduct)
		r.Post("/product", h.AddProduct)
		r.Put("/product/{code}", h.UpdateProduct)
		r.Delete("/product/{code}", h.DeleteProduct)

		//Promotion
		r.Get("/promotion", h.GetAllPromotion)
		r.Get("/promotion/{code}", h.GetDetailPromotion)
		r.Post("/promotion", h.AddPromotion)
		r.Put("/promotion/{code}", h.UpdatePromotion)
		r.Delete("/promotion/{code}", h.DeletePromotion)

		//Category
		r.Get("/category", h.GetAllCategory)
		r.Post("/category", h.AddCategory)
		r.Delete("/category/{code}", h.DeleteCategory)

		//Export
		r.Get("/export", h.SearchExport)

		//Download
		r.Get("/download", h.GetAllDownload)
		r.Post("/download", h.AddDownload)
		r.Delete("/download/{code}", h.DeleteDownload)
	})
}
