package main

import (
	"net/http"

	"github.com/manuvirk/bookings/pkg/config"
	"github.com/manuvirk/bookings/pkg/handlers"

	"github.com/go-chi/chi"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	//mux.Use(middleware.Recoverer)
	mux.Use(NoSruf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	//mux.Use(WriteToConsole)
	return mux
}
