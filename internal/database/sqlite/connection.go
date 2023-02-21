package sqlite

import (
	"database/sql"
	"github.com/maxicapodacqua/nearby/internal/config"
	"log"
	_ "modernc.org/sqlite"
	"os"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite", os.Getenv(config.SQLiteDNS))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
