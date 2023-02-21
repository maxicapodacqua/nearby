package router

import (
	"encoding/json"
	"net/http"
)

type PingResponse struct {
	Pong string `json:"pong"`
}

// Ping
// Returns success as long as the web API is running
func Ping() (path string, handler HandlerFunc) {
	return "/ping", func(w http.ResponseWriter, r *http.Request) error {
		resp := NewResponse(PingResponse{Pong: "success"})

		rMarshal, err := json.Marshal(resp)
		if err != nil {
			return err
		}

		if _, err = w.Write(rMarshal); err != nil {
			return err
		}
		return nil
	}
}
