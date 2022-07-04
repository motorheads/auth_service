package api

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/motorheads/auth_service/storage"
)

func Register(c *fiber.Ctx) error {
	var user map[string]string

	if err := c.BodyParser(&user); err != nil {
		fmt.Println("Error while parsing body")
		fmt.Println(err)
		return err
	}

	err := storage.CreateUser(user)
	if err != nil {
		fmt.Println("Erro while creating user in the database")
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK)
}

func Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func User(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
