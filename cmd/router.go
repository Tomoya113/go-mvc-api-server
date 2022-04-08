package main

import (
	"github.com/Tomoya113/go-mvc-api-server/pkg/controller"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitializeRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", controller.GetUsers)
		r.Post("/", controller.CreateUser)
	})

	return r
}
