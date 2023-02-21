package sqlite

import (
	"database/sql"
	"github.com/maxicapodacqua/nearby/internal/config"
	"log"
	_ "modernc.org/sqlite"
	"os"
)

func Connect() *sql.DB {
	log.Printf("Connecting to sqlite database\n")
	db, err := sql.Open("sqlite", os.Getenv(config.SQLiteDNS))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
