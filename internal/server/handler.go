package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func ping() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		resp := GeneralResponse[PingResponse]{
			Data: PingResponse{
				Pong: "success",
			},
		}

		rMarshal, err := json.Marshal(resp)
		if err != nil {
			log.Printf("Something went wrong in /ping %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(rMarshal)
	})
}
