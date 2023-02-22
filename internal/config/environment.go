// Package config
// Manages general app config, mostly based in environment
// variables and holds constants to access them. To access config variables in
// this project you have to use os.GetEnv
package config

import (
	"log"
	"os"
)

const (
	Port        = "PORT"
	DefaultPort = "8080"

	SQLiteDNS        = "SQLITE_DNS"
	DefaultSQLiteDNS = "data/db.sqlite"
)

func setDefault(key, val string) {
	log.Printf("%v not set, using default value of %v", key, val)
	if err := os.Setenv(key, val); err != nil {
		log.Fatalf("Something went wrong changing env var %v", err)
	}
}

// SetWithDefaults
// Initializes the configuration of the project as environment variables
// All config vars are accessed using os.GetEnv
func SetWithDefaults() {
	if _, ok := os.LookupEnv(Port); !ok {
		setDefault(Port, DefaultPort)
	}
	if _, ok := os.LookupEnv(SQLiteDNS); !ok {
		setDefault(SQLiteDNS, DefaultSQLiteDNS)
	}
}
