package models

type Site struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Url  string `db:"url"`
}
