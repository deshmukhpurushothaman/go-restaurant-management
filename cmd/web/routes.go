package main

import (
	"net/http"

	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/handlers"
	middlewares "github.com/deshmukhpurushothaman/go-restaurant-management/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Heartbeat("/"))

	mux.Route("/auth", func(r chi.Router) {
		r.Post("/register", handlers.Repo.RegisterUser)
		r.Post("/login", handlers.Repo.LoginHandler)
	})

	mux.Route("/category", func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.With(middlewares.RoleMiddleware(middlewares.RoleAdmin)).Post("/create", handlers.Repo.CreateCategory)
		r.Get("/all", handlers.Repo.GetCategories)
		r.With(middlewares.RoleMiddleware(middlewares.RoleAdmin)).Put("/update", handlers.Repo.UpdateCategory)
		r.With(middlewares.RoleMiddleware(middlewares.RoleAdmin)).Delete("/{id}", handlers.Repo.DeleteCatogory)
		r.Get("/{id}", handlers.Repo.GetCategoryById)
	})

	mux.Route("/food", func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.With(middlewares.RoleMiddleware(middlewares.RoleAdmin)).Post("/create", handlers.Repo.CreateFood)
		r.Get("/all", handlers.Repo.GetAllFoods)
		r.With(middlewares.RoleMiddleware(middlewares.RoleAdmin)).Put("/update", handlers.Repo.UpdateFood)
		r.With(middlewares.RoleMiddleware(middlewares.RoleAdmin)).Delete("/{id}", handlers.Repo.DeleteFood)
		r.Get("/{id}", handlers.Repo.GetFoodById)
	})
	return mux
}
