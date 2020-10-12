package main

// import packages from standard lib
import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the entrypoint to the application. It starts a web server, listening on port 8080,
// and sets up two simple routes.
func main() {
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/about", AboutPageHandler)
	http.HandleFunc("/contact", ContactPageHandler)

	log.Println(fmt.Sprintf("Starting web server on port %s", portNumber))
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server : ", err)
		return
	}
}
