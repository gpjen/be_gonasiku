package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gonasiku.com/db"
	"gonasiku.com/user"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	db.InitPostgresDB()

	userRepository := user.NewRepository(db.PostgresDB)

	// user, _ := userRepository.Create(user.UserRegister{
	// 	Name:     "test",
	// 	Email:    "test@one.com",
	// 	Password: "rahasiatest",
	// 	Phone:    "0862516522222",
	// 	Addres:   "cimahi",
	// 	Gender:   "p",
	// })

	// user, _ := userRepository.FindByID(1)

	// user, _ := userRepository.Update(3, user.UserUpdate{
	// 	Name:   "ariana",
	// 	Email:  "ariana@grenade.com",
	// 	Phone:  "01212121212122",
	// 	Addres: "pacitan",
	// 	Gender: "p",
	// })

	// user, _ := userRepository.Delete(3)

	user, _ := userRepository.DeletePermanent(3)

	fmt.Println(user)

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
