package controlers

import (
	"encoding/json"
	"net/http"
	"vinimpv/gogums/models"
	"vinimpv/gogums/services"
)

type RepositorySerializer struct {
	Id     int64  `json:"id,omitempty"`
	SiteId int64  `json:"site_id"`
	Url    string `json:"url"`
	Key    string `json:"key"` //ssh key as plain text
}

func (rs *RepositorySerializer) validate() error {
	return nil
}

func (rs *RepositorySerializer) save() (models.Repository, error) {
	return services.Repositories.CreateRepository(rs.Url, rs.Key, rs.SiteId)
}

func GetRepositories(w http.ResponseWriter, r *http.Request) {
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

func CreateRepository(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rs := &RepositorySerializer{}
	json.NewDecoder(r.Body).Decode(rs)
	// check for validation error
	err := rs.validate()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	repository, err := rs.save()

	b, err := json.Marshal(repository)
	if err != nil {
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
