package main

// import packages from standard lib
import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
	"tsawler/go-course/pkg/config"
	"tsawler/go-course/pkg/driver"
	"tsawler/go-course/pkg/handlers"
	"tsawler/go-course/pkg/templates"
)

const portNumber = ":8080"
const inProduction = false

var session *scs.SessionManager
var dbUser = "dbUser"
var dbPass = "verysecret"
var dbHost = "127.0.0.1"
var dbPort = "3306"
var databaseName = "myapp"
var dbSsl = "false"
var databaseEngine = "mysql"

// main is the entry point to the application. It starts a web server, listening on port 8080,
// connects to the database, sets up sessions, and passes in our routes file
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

	// connect to database
	dsnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=%s&collation=utf8_unicode_ci&timeout=5s&readTimeout5s", dbUser, dbPass, dbHost, dbPort, databaseName, dbSsl)

	log.Printf("Connecting to database %s: %s & initializing pool....", databaseEngine, databaseName)
	var db *driver.DB
	db, err := driver.ConnectSQL(dsnString)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := db.SQL.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// init template cache
	err = templates.NewTemplateCache(&app)
	if err != nil {
		log.Fatal(err)
	}

	// create database repository
	repo := handlers.NewDatabaseRepo(db, &app)

	// give the app config and database to handler functions
	handlers.NewHandlers(repo)

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
