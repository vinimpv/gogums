package controlers

import (
	"encoding/json"
	"net/http"
	"vinimpv/gogums/models"
	"vinimpv/gogums/services"
)

type ResourcesDefinitionSerializer struct{}

type ResourceSerializer struct {
	GroupName string           `json:"group_name"`
	Resource  *models.Resource `json:"resource"`
}

type ResourcesGroupSerializer struct {
	Name        string
	Description string
	FolderPath  string
	Main        *ResourcesDefinitionSerializer
	List        []ResourcesDefinitionSerializer
}

// opens up site repo dir
// parses the templates
// iterate through the names, read files in folder convert to Resource

func ResourceControler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	site := r.Context().Value("site").(*models.Site)
	rs := &ResourceSerializer{}
	json.NewDecoder(r.Body).Decode(rs)
	resource, err := services.SaveResource(site, rs.Resource, rs.GroupName)
	responseData, err := json.Marshal(resource)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
