package main

import (
	"net/http"
	"github.com/countBreadedDice/booking_go_try/pkg/config"
	"github.com/countBreadedDice/booking_go_try/pkg/handlers"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	// mux.Get("/sum", http.HandlerFunc(handlers.Repo.Sum))
	// mux.Get("/divide", http.HandlerFunc(handlers.Repo.Divide))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/sum", http.HandlerFunc(handlers.Repo.Sum))
	mux.Get("/divide", http.HandlerFunc(handlers.Repo.Divide))
	return mux
}
