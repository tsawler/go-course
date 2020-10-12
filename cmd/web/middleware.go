package main

import (
	"fmt"
	"net/http"
	"tsawler/go-course/pkg/helpers"
)

// RecoverPanic recovers from a panic
func RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// Check if there has been a panic
			if err := recover(); err != nil {
				// return a 500 Internal Server response
				helpers.ServerError(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
