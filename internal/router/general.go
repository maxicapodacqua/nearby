// Package router
// Adds different routes definitions for the API
// including responses and payloads
package router

import "net/http"

// HandlerFunc
// Common function to replace http.HandlerFunc
// We return an error to bubble up the failure to a top level middleware
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

type GeneralResponse[T any] struct {
	Data T `json:"data"`
}
