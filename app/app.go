package app

import (
	"net/http"

	"github.com/go-chi/chi"
)

var router *chi.Mux

func init() {
	router = chi.NewRouter()
}

func StartApp() {
	http.ListenAndServe(":3000", router)
}
