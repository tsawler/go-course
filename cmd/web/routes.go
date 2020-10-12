package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"tsawler/go-course/pkg/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(RecoverPanic)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Get("/", handlers.HomePageHandler)
	mux.Get("/about", handlers.AboutPageHandler)
	mux.Get("/contact", handlers.ContactPageHandler)

	return mux
}
