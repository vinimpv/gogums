package controlers

import (
	"encoding/json"
	"net/http"
	"vinimpv/gogums/models"
	"vinimpv/gogums/services"
)

type SiteSerializer struct {
	Id         int                        `json:"id,omitempty"`
	Name       string                     `json:"name"`
	Url        string                     `json:"url"`
	Repository RepositorySerializer       `json:"repository"`
	Resources  []ResourcesGroupSerializer `json:"resources"`
}

func (ss *SiteSerializer) fromSite(site models.Site) {

}

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
