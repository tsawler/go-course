package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func routes() http.Handler {
	mux := pat.New()

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	mux.Get("/", http.HandlerFunc(HomePageHandler))
	mux.Get("/about", http.HandlerFunc(AboutPageHandler))
	mux.Get("/contact", http.HandlerFunc(ContactPageHandler))

	return mux
}
