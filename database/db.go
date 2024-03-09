package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"lobinv-server/config"
)

var DB *sql.DB

func ConnectDB(c config.Config) {
	db, err := sql.Open("sqlite3", c.Database.Uri)

	if err != nil {
		panic(err)
	}

	DB = db
}
