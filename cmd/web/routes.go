package main

import (
	"net/http"

	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Heartbeat("/"))

	mux.Route("/category", func(mux chi.Router) {
		mux.Post("/create", handlers.Repo.CreateCategory)
		mux.Get("/all", handlers.Repo.GetCategories)
		mux.Put("/update", handlers.Repo.UpdateCategory)
		mux.Delete("/delete/{id}", handlers.Repo.DeleteCatogory)
		mux.Get("/{id}", handlers.Repo.GetCategoryById)
	})
	return mux
}
