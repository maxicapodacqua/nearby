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
// @Summary Checks if the API is healthy
// @Produce json
// @Tags			system
// @Accept			json
// @Success 200 {object} GeneralResponse[HealthResponse] "API status"
// @Failure	500	{object} GeneralResponse[HealthResponse] "Server error, likely database is down"
// @Router			/health [get]
func Health(db *sql.DB) (path string, handler HandlerFunc) {
	return "/health", func(w http.ResponseWriter, r *http.Request) error {
		resp := NewResponse(HealthResponse{Database: "healthy"})
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
