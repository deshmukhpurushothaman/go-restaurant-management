package main

import (
	"net/http"

	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Heartbeat("/ping"))
	// mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte("Hello World"))
	// })

	mux.Get("/", handlers.Repo.DummyTest)
	return mux
}
