package main

import (
	"net/http"
	"web3/pkg/config"
	"web3/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(LogRequestInfo)

	mux.Use(NoSurf)
	mux.Use(SetupSession)

	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	mux.Get("/login", handlers.Repo.LoginHandler)
	mux.Get("/page", handlers.Repo.PageHandler)
	mux.Get("/makepost", handlers.Repo.MakePostHandler)

	mux.Post("/makepost", handlers.Repo.PostMakePostHandler)

	// Returns a handler that provides access to our
	// static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
