package http

import (
	"github.com/VyacheslavKuzharov/gophermart/internal/di"
	authhandler "github.com/VyacheslavKuzharov/gophermart/internal/transport/http/handlers/auth"
	"github.com/VyacheslavKuzharov/gophermart/internal/transport/http/middlewares"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router *chi.Mux, container *di.Container) {
	router.Use(middlewares.Logger(container.Logger))

	// api routes
	router.Route("/api", func(api chi.Router) {
		// user routes
		api.Route("/user", func(user chi.Router) {
			user.Group(func(public chi.Router) {
				handler := authhandler.New(container.GetAuthUseCase())

				public.Post("/register", handler.SignUp)
				public.Post("/login", handler.SignIn)
			})
			user.Group(func(private chi.Router) {
				//private.Use(middlewares.Auth)

				//private.Get("/orders", ordersHandler)
			})
		})
	})
}
