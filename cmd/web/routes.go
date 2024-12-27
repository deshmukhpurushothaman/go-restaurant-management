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

	mux.Route("/auth", func(r chi.Router) {
		r.Post("/register", handlers.Repo.RegisterUser)
	})

	mux.Route("/category", func(r chi.Router) {
		r.Post("/create", handlers.Repo.CreateCategory)
		r.Get("/all", handlers.Repo.GetCategories)
		r.Put("/update", handlers.Repo.UpdateCategory)
		r.Delete("/{id}", handlers.Repo.DeleteCatogory)
		r.Get("/{id}", handlers.Repo.GetCategoryById)
	})

	mux.Route("/food", func(r chi.Router) {
		r.Post("/create", handlers.Repo.CreateFood)
		r.Get("/all", handlers.Repo.GetAllFoods)
		r.Put("/update", handlers.Repo.UpdateFood)
		r.Delete("/{id}", handlers.Repo.DeleteFood)
		r.Get("/{id}", handlers.Repo.GetFoodById)
	})
	return mux
}
