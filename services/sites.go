package services

import (
	"fmt"
	"vinimpv/gogums/db"
	"vinimpv/gogums/models"
)

type sitesService struct{}

type SitesServiceInterface interface {
	GetSites() ([]models.Site, error)
	GetSite(siteId int64) (*models.Site, error)
	CreateSite(name string, url string) (models.Site, error)
}

var Sites SitesServiceInterface

func init() {
	Sites = &sitesService{}

}

func (ss *sitesService) GetSites() ([]models.Site, error) {
	sites := []models.Site{}
	//TODO fix scan of nonexistent repository
	sql := `SELECT
		st.id, st.name, st.url,
		rp.id "repository.id", rp.url "repository.url", rp.key "repository.key", rp.site_id "repository.site_id"
		FROM sites st 
		left JOIN repositories rp
		on st.id = rp.site_id`

	err := db.DB.Select(&sites, sql)

	if err != nil {
		fmt.Println(err)
	}
	return sites, nil
}

func (ss *sitesService) GetSite(siteId int64) (*models.Site, error) {
	site := &models.Site{}
	sql := `SELECT
		st.id, st.name, st.url,
		rp.id "repository.id", rp.url "repository.url", rp.key "repository.key", rp.site_id "repository.site_id"
		FROM sites st 
		INNER JOIN repositories rp
		on st.id = rp.site_id
		WHERE st.id = ?`

	err := db.DB.Get(site, sql, siteId)
	return site, err
}

func (ss *sitesService) CreateSite(name string, url string) (models.Site, error) {
	result := db.DB.MustExec("INSERT INTO sites (name, url) VALUES ($1, $2)", name, url)
	id, err := result.LastInsertId()
	if err != nil {
		return models.Site{}, err
	}
	return models.Site{Id: id, Name: name, Url: url}, nil
}
