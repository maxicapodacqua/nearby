package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type GeneralResponse[T any] struct {
	Data T `json:"data"`
}
type PingResponse struct {
	Pong string `json:"pong"`
}

func main() {
	log.Printf("Starting server \n")
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
	
	log.Printf("Server started\n")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
