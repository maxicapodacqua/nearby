package router

import (
	"errors"
	"net/http"
)

type HealthResponse struct {
	Database string `json:"database"`
}

// Health
// Checks if API connections are okay
func Health() (path string, handler HandlerFunc) {
	return "/health", func(w http.ResponseWriter, r *http.Request) error {
		w.WriteHeader(500)
		return errors.New("unimplemented")
	}
}
