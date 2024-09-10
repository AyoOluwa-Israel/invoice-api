package routes

import (
	"github.com/AyoOluwa-Israel/invoice-api/controllers"
	"github.com/gofiber/fiber/v2"
)


func UserRoutes(app fiber.Router){
	app.Get("/users", controllers.GetAllUsers)
	app.Get("/user/:user_id", controllers.GetUser)
	app.Post("/user", controllers.RegisterUser)
}