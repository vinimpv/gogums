package app

import "vinimpv/gogums/controlers"

func mapUrls() {
	router.Get("/sites", controlers.GetSites)
	router.Post("/sites", controlers.CreateSite)
}
