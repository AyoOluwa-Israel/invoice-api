package routes

import (
	"github.com/AyoOluwa-Israel/invoice-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func PaymentInformationRoutes(app fiber.Router){
	app.Post("/user/payment", controllers.CreatePaymentInfo)
	
}
