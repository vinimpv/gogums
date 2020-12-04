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

func (ss *SiteSerializer) validate() error {
	return nil
}

func (ss *SiteSerializer) save() (models.Site, error) {
	// check if id exists and create/update
	return services.Sites.CreateSite(ss.Name, ss.Url)
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

func GetSite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	site := r.Context().Value("site").(*models.Site)
	site.Repository.Clone()
	a, _ := services.ParseResources(site)
	site.Resources = a
	b, err := json.Marshal(site)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func CreateSite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ss := &SiteSerializer{}
	json.NewDecoder(r.Body).Decode(ss)
	// check for validation error
	err := ss.validate()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	site, err := ss.save()

	b, err := json.Marshal(site)
	if err != nil {
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
