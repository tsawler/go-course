package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"log"
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

// SessionLoad loads the session on requests
func SessionLoad(next http.Handler) http.Handler {
	log.Println("Loading session...")
	return session.LoadAndSave(next)
}

// NoSurf implements CSRF protection
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   inProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
