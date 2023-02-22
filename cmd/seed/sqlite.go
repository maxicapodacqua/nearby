// Creates the needed tables for the sqlite database
// and inserts data
package main

import (
	"github.com/maxicapodacqua/nearby/internal/config"
	"github.com/maxicapodacqua/nearby/internal/database/sqlite"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func main() {

	log.Println("Starting seed process")
	config.SetWithDefaults()
	db := sqlite.Connect()

	log.Println("Ping database")
	check(db.Ping())

	log.Println("Creating table")
	schemaF, err := os.ReadFile("data/schema.sql")
	check(err)

	_, err = db.Exec(string(schemaF))
	check(err)

	log.Println("Inserting data in fresh table")
	dataF, err := os.ReadFile("data/data.sql")
	check(err)
	r, err := db.Exec(string(dataF))
	check(err)
	rowsAf, _ := r.RowsAffected()
	log.Printf("Rows affected: %v", rowsAf)

	var count int
	err = db.QueryRow("SELECT count(*) FROM `nearby_area_codes`").Scan(&count)
	check(err)

	log.Printf("Return count from select: %v\n", count)

	log.Println("End seed process")
}
