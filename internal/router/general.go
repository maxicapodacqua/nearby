// Package router
// Adds different routes definitions for the API
// including responses and payloads
package router

import "net/http"

// HandlerFunc
// Common function to replace http.HandlerFunc
// We return an error to bubble up the failure to a top level middleware
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

type GeneralResponse[T interface{}] struct {
	Data  T       `json:"data"`
	Error *string `json:"error"`
}

func NewResponse[T interface{}](dataStruct T) GeneralResponse[T] {
	return GeneralResponse[T]{Data: dataStruct}
}

func NewErrorResponse(err error) GeneralResponse[interface{}] {
	str := err.Error()
	return GeneralResponse[interface{}]{
		Data:  nil,
		Error: &str,
	}
}
