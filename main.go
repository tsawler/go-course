package main

import (
	"fmt"
	"net/http"
)

// main starts a simple web server on port 8080, and writes hello world to the browser in two languages.
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Hello, world! こんにちは世界")
	})

	_ = http.ListenAndServe(":8080", nil)
}
