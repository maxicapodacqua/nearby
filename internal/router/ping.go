package router

import (
	"encoding/json"
	"net/http"
)

type PingResponse struct {
	Pong string `json:"pong"`
}

// Ping
// @Summary Returns success as long as the web API is running
// @Produce json
// @Tags			system
// @Accept			json
// @Success 200 {object} GeneralResponse[PingResponse] "Pong if the API is running"
// @Failure	500	{object} GeneralResponse[string] "Server error"
// @Router			/ping [get]
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
