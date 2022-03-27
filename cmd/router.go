package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-mvc-api-server/pkg/controller"
	"go-mvc-api-server/pkg/model"
	"go-mvc-api-server/pkg/view"
)

func InitializeRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	v := view.NewUserView()
	m := model.NewUserModel()
	c := controller.NewUserController(v, m)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", c.GetUsers)
		r.Post("/", c.CreateUser)
	})

	return r
}
