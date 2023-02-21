package router

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Database string `json:"database"`
}

// Health
// Checks if API connections are okay
func Health(db *sql.DB) (path string, handler HandlerFunc) {
	return "/health", func(w http.ResponseWriter, r *http.Request) error {
		resp := GeneralResponse[HealthResponse]{Data: HealthResponse{Database: "healthy"}}
		dbErr := db.PingContext(r.Context())
		if dbErr != nil {
			w.WriteHeader(500)
			resp.Data.Database = "unhealthy"
		}
		rMarshal, _ := json.Marshal(resp)
		_, _ = w.Write(rMarshal)

		return dbErr
	}
}
