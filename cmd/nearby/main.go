package main

import (
	"github.com/maxicapodacqua/nearby/internal/config"
	"github.com/maxicapodacqua/nearby/internal/server"
	"log"
)

//	@title			Nearby API
//	@version		1.0
//	@description	Nearby API, to search nearby phone area codes
//
// @host		localhost:8080
// @BasePath	/
func main() {
	// Adds microseconds to logger
	log.Default().SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	config.SetWithDefaults()
	server.Start()
}
