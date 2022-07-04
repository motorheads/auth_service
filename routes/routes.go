package routes

import "github.com/gofiber/fiber/v2"

func New() *fiber.App {
	r := fiber.New()

	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Test")
	})

	return r
}
