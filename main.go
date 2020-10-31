package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"vinimpv/gogums/services"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	site, err := services.Sites.CreateSite("lalala", "sadfhjash")
	repo, err := services.Repositories.CreateRepository("www.google.com", "minha_chave", site.Id)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(site)
	fmt.Println(repo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		sites, _ := services.Sites.GetSites()
		b, err := json.Marshal(sites)
		if err != nil {
			http.Error(w, err.Error(), 422)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})
	http.ListenAndServe(":3000", r)
}
