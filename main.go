package main

// import packages from standard lib
import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

// Home handles requests to home page
func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello, world!")
}

// main is the entrypoint to the application. It starts a web server, listening on port 8080,
// and sets up two simple routes.
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", AboutPageHandler)
	http.HandleFunc("/contact", ContactPageHandler)

	log.Println(fmt.Sprintf("Starting web server on port %s", portNumber))
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server : ", err)
		return
	}
}
