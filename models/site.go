package models

type Site struct {
	Id        int64      `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	Url       string     `db:"url" json:"url"`
	Repo      Repository `db:"repository"`
	Resources ResourcesGroup
}
