package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gonasiku.com/db"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	db.InitMysqlDB()

	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "http://localhost:3000",
			AllowMethods: "GET, POST, PUT, DELETE",
			AllowHeaders: "Origin, Content-Type, Accept",
		},
	))

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
