package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(HomePageHandler))
	mux.Get("/about", http.HandlerFunc(AboutPageHandler))
	mux.Get("/contact", http.HandlerFunc(ContactPageHandler))

	return mux
}
