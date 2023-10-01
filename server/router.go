package main

import (
	"server/handlers"
	"server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func MainRouter(r fiber.Router) {

	mainRoutes := r.Group("/")
	mainRoutes.Post("/auth", handlers.AuthHandler)
	mainRoutes.Post("/login", handlers.LoginHandler)
	mainRoutes.Post("/signup", handlers.SignUpHandler)

	authRoutes := r.Group("/", middlewares.AuthMiddleware)
	authRoutes.Post("/updateuser", handlers.UpdateUserHandler)
	authRoutes.Delete("/deleteuser", handlers.DeleteUserHandler)
}
