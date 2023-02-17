// Package server
// Defines configuration and starting point
// for the Web API
package server

import (
	"github.com/maxicapodacqua/nearby/internal/router"
	"log"
	"net/http"
	"time"
)

func configureRoute(path string, handler router.HandlerFunc) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := handler(w, r)
		if err != nil {
			log.Printf("ERROR [%v] %v %v ", r.Method, r.URL, err)
			return
		}
		log.Printf("INFO [%v] %v", r.Method, r.URL)
	})
}

// Start
// Initializes the Web API with all its routes
func Start() {
	// Adds microseconds to logger
	log.Default().SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Printf("Starting server \n")

	configureRoute(router.Ping())
	configureRoute(router.Health())

	log.Printf("Server started\n")
	s := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		ErrorLog:       log.Default(),
	}
	log.Fatal(s.ListenAndServe())
}
