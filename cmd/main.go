package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	r := fiber.New()

	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Test")
	})

	r.Listen(":8000")
}
