package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

var schema = `
CREATE TABLE IF NOT EXISTS sites (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(80) DEFAULT '',
	url VARCHAR(200) DEFAULT ''
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
