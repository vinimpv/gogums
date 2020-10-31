package controlers

import (
	"encoding/json"
	"net/http"
	"vinimpv/gogums/services"
)

type RepositorySerializer struct {
	Id  int    `json:"id,omitempty"`
	Url string `json:"url"`
	Key string `json:"key"` //ssh key as plain text
}

func GetRepository(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sites, _ := services.Repositories.GetRepositories()
	b, err := json.Marshal(sites)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
