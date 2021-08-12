package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"example.com/hello/database"
	"example.com/hello/routes"
)

//using fiber framework
func main() {
	database.Connect()

	app := fiber.New()

// this will help frontend to receive cookies from backend
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	app.Listen(":8000")
}




