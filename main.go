package main

// import packages from standard lib
import (
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

// main is the entrypoint to the application. It starts a web server, listening on port 8080,
// and sets up two simple routes.
func main() {
	srv := &http.Server{
		Addr:              portNumber,
		Handler:           routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	log.Printf("Starting HTTP server on port %s....", portNumber)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
