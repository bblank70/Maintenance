package main

import (
	"net/http"

	"github.com/bblank70/bookings/pkg/config"
	"github.com/bblank70/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	// mux := pat.New()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	//forms
	mux.Get("/pump-form", handlers.Repo.Pump)
	mux.Post("/pump-form", handlers.Repo.PostPump)
	mux.Get("/stage", handlers.Repo.Stage)
	mux.Get("/repaired", handlers.Repo.Repaired)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
