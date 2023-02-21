package router

import (
	"errors"
	"github.com/maxicapodacqua/nearby/internal/database/sqlite"
	"log"
	"net/http"
)

type HealthResponse struct {
	Database string `json:"database"`
}

// Health
// Checks if API connections are okay
func Health() (path string, handler HandlerFunc) {
	return "/health", func(w http.ResponseWriter, r *http.Request) error {

		db := sqlite.Connect()
		defer db.Close()
		log.Printf("DB connection: %+v", db)

		w.WriteHeader(500)
		return errors.New("unimplemented")
	}
}
