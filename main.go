package main

// import packages from standard lib
import (
	"fmt"
	"log"
	"net/http"
)

// Home handles requests to home page
func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello, world!")
}

// About handles requests to the about page
func About(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "You chose the about page!")
}

// main is the entrypoint to the application. It starts a web server, listening on port 8080,
// and sets up two simple routes.
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server : ", err)
		return
	}
}
