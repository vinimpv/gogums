package services

import (
	"fmt"
	"vinimpv/gogums/db"
	"vinimpv/gogums/models"
)

type sitesService struct{}

type SitesServiceInterface interface {
	GetSites() ([]models.Site, error)
	CreateSite(name string, url string) (models.Site, error)
}

var Sites SitesServiceInterface

func init() {
	Sites = &sitesService{}

}

func (ss *sitesService) GetSites() ([]models.Site, error) {
	sites := []models.Site{}
	err := db.DB.Select(&sites, "SELECT * FROM sites")
	if err != nil {
		fmt.Println(err)
	}
	return sites, nil
}

func (ss *sitesService) CreateSite(name string, url string) (models.Site, error) {
	result := db.DB.MustExec("INSERT INTO sites (name, url) VALUES ($1, $2)", name, url)
	id, err := result.LastInsertId()
	if err != nil {
		return models.Site{}, err
	}
	return models.Site{Id: id, Name: name, Url: url}, nil
}
