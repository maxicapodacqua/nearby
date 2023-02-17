// Package configuration
// Manages general app configuration, mostly based in environment variables
// and holds constants to access them
package configuration

import (
	"log"
	"os"
)

const (
	Port        = "PORT"
	DefaultPort = "8080"
)

func setDefault(key, val string) {
	log.Printf("%v not set, using default value of %v", key, val)
	if err := os.Setenv(key, val); err != nil {
		log.Fatalf("Something went wrong changing env var %v", err)
	}
}

func SetWithDefaults() {
	if _, ok := os.LookupEnv(Port); !ok {
		setDefault(Port, DefaultPort)
	}
}
