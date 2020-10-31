package app

import (
	"vinimpv/gogums/controlers"

	"github.com/go-chi/chi"
)

func mapUrls() {
	router.Route("/sites", func(r chi.Router) {
		r.Get("/", controlers.GetSites)
		r.Post("/", controlers.CreateSite)
	})
	router.Route("/repositories", func(r chi.Router) {
		r.Get("/", controlers.GetRepository)
		r.Post("/", controlers.CreateRepository)
	})
}
