package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

var schema = `
DROP TABLE sites;
DROP TABLE repositories;
CREATE TABLE IF NOT EXISTS sites (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(80) DEFAULT '',
	url VARCHAR(200) DEFAULT ''
);
CREATE TABLE IF NOT EXISTS repositories (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	url VARCHAR(200) DEFAULT '',
	key VARCHAR(200) DEFAULT '',
	site_id INTEGER
);
`

func init() {
	db, err := sqlx.Connect("sqlite3", "./gogums.db")
	db.MustExec(schema)
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
}
