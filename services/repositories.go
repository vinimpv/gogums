package services

import (
	"fmt"
	"vinimpv/gogums/db"
	"vinimpv/gogums/models"
)

type repositoriesService struct{}

type RepositoriesServiceInterface interface {
	GetRepositories() ([]models.Repository, error)
	CreateRepository(url, key string, siteId int64) (models.Repository, error)
}

var Repositories RepositoriesServiceInterface

func init() {
	Repositories = &repositoriesService{}

}

func (rs *repositoriesService) GetRepositories() ([]models.Repository, error) {
	sites := []models.Repository{}
	err := db.DB.Select(&sites, "SELECT * FROM repositories")
	if err != nil {
		fmt.Println(err)
	}
	return sites, nil
}

func (rs *repositoriesService) CreateRepository(url, key string, siteId int64) (models.Repository, error) {
	result := db.DB.MustExec("INSERT INTO repositories (url, key, site_id) VALUES ($1, $2, $3)", url, key, siteId)
	id, err := result.LastInsertId()
	if err != nil {
		return models.Repository{}, err
	}
	return models.Repository{Id: id, Url: url, Key: key, SiteId: siteId}, nil
}
