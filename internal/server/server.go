// Package server
// Defines config and starting point
// for the Web API
package server

import (
	"github.com/maxicapodacqua/nearby/internal/config"
	"github.com/maxicapodacqua/nearby/internal/database/sqlite"
	"github.com/maxicapodacqua/nearby/internal/router"
	"log"
	"net/http"
	"os"
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
	log.Printf("Starting server\n")

	db := sqlite.Connect()

	configureRoute(router.Ping())
	configureRoute(router.Health(db))
	configureRoute(router.Nearby(db))

	s := &http.Server{
		Addr:           ":" + os.Getenv(config.Port),
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		ErrorLog:       log.Default(),
	}
	log.Printf("Server started at %v \n", s.Addr)
	log.Fatal(s.ListenAndServe())
}
