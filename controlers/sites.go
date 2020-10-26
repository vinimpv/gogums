package controlers

import (
	"encoding/json"
	"net/http"
	"vinimpv/gogums/services"
)

func GetSites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sites, _ := services.Sites.GetSites()
	b, err := json.Marshal(sites)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func CreateSite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sites, _ := services.Sites.GetSites()
	b, err := json.Marshal(sites)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
