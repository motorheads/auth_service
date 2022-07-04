package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"github.com/motorheads/auth_service/config"
	"github.com/motorheads/auth_service/routes"
)

func main() {
	db, err := sql.Open("postgres", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Error while connecting to the database")
		fmt.Println(err)
	}
	config.DB = db

	defer config.DB.Close()

	router := routes.New()

	router.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	router.Listen(":8081")
}
