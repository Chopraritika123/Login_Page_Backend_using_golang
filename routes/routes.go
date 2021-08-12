package routes

import (
	"github.com/gofiber/fiber/v2"

	"example.com/hello/controllers"
)

func Setup(app *fiber.App){
	app.Post("/api/register", controllers.Register)
	app.Post("/api/loginUser", controllers.LoginUser)
	app.Get("/api/User", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}
