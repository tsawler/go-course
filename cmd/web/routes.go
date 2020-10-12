package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"tsawler/go-course/pkg/config"
	"tsawler/go-course/pkg/handlers"
)

func routes(app config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(RecoverPanic)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Get("/", handlers.HomePageHandler(app))
	mux.Get("/about", handlers.AboutPageHandler(app))
	mux.Get("/contact", handlers.ContactPageHandler(app))

	return mux
}
