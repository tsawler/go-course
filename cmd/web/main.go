package main

// import packages from standard lib
import (
	"log"
	"net/http"
	"time"
	"tsawler/go-course/pkg/config"
	"tsawler/go-course/pkg/templates"
)

const portNumber = ":8080"

// main is the entrypoint to the application. It starts a web server, listening on port 8080,
// and passes it our routes file
func main() {
	var app config.AppConfig
	app.UseCache = false

	// init template cache
	err := templates.NewTemplateCache(&app)
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:              portNumber,
		Handler:           routes(app),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	log.Printf("Starting HTTP server on port %s....", portNumber)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
