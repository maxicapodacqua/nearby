package sqlite

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite", "data/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
