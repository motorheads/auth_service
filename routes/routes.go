package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/motorheads/auth_service/api"
)

func New() *fiber.App {
	router := fiber.New()

	router.Group("/api")

	router.Post("/api/register", api.Register)
	router.Post("/api/login", api.Login)
	router.Get("/api/user", api.User)
	router.Post("/api/logout", api.Logout)

	return router
}
