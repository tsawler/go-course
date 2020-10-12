package main

// import packages from standard lib
import (
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
	"tsawler/go-course/pkg/config"
	"tsawler/go-course/pkg/handlers"
	"tsawler/go-course/pkg/templates"
)

const portNumber = ":8080"
const inProduction = false

var session *scs.SessionManager

// main is the entry point to the application. It starts a web server, listening on port 8080,
// and passes it our routes file
func main() {
	var app config.AppConfig
	app.UseCache = false

	// set up session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = inProduction

	// put the session in app config
	app.Session = session

	// init template cache
	err := templates.NewTemplateCache(&app)
	if err != nil {
		log.Fatal(err)
	}

	// give the app config to all handler functions, and not just
	// actual handlers
	handlers.NewHandlers(&app)

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
