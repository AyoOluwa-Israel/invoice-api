package routes

import (
	"github.com/AyoOluwa-Israel/invoice-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func WebsiteRoutes(app fiber.Router) {
	app.Post("/submit-request", controllers.PostMessage)
}
