package models

type Site struct {
	Id         int64          `db:"id" json:"id"`
	Name       string         `db:"name" json:"name"`
	Url        string         `db:"url" json:"url"`
	Repository Repository     `db:"repository" json:"repository"`
	Resources  *SiteResources `json:"resources"`
}
