package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"tsawler/go-course/pkg/config"
	"tsawler/go-course/pkg/handlers"
)

func routes(app config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// default middleware
	mux.Use(RecoverPanic)
	mux.Use(SessionLoad)
	mux.Use(NoSurf)

	// static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// application routes
	mux.Get("/", handlers.Repo.HomePageHandler(app))
	mux.Get("/about", handlers.Repo.AboutPageHandler(app))
	mux.Get("/contact", handlers.Repo.ContactPageHandler(app))
	mux.Post("/contact", handlers.Repo.PostContactPageHandler(app))

	return mux
}
