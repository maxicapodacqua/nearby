package main

import (
	"github.com/maxicapodacqua/nearby/internal/configuration"
	"github.com/maxicapodacqua/nearby/internal/server"
	"log"
)

func main() {
	// Adds microseconds to logger
	log.Default().SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	configuration.SetWithDefaults()
	server.Start()
}
