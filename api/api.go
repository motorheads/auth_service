package api

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/motorheads/auth_service/config"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Error while parsing body")
		fmt.Println(err)
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	query := `
		INSERT INTO users(
			name,
			email,
			password
		) VALUES (
			$1,
			$2,
			$3
		);`
	_, err := config.DB.Exec(query, data["name"], data["email"], password)
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
