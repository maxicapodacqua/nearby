package server

import (
	"log"
	"net/http"
)

func Start() {
	log.Printf("Starting server \n")

	ping()

	log.Printf("Server started\n")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
