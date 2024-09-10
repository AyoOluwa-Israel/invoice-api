package routes

import (
	"github.com/AyoOluwa-Israel/invoice-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func InvoiceRoutes(app fiber.Router) {

	app.Post("/invoice", controllers.CreateInvoice)
	app.Get("/invoice", controllers.GetAllInvoice)

	app.Get("/invoice/:invoice_id", controllers.GetInvoiceByID)
	app.Put("/invoice/:invoice_id", controllers.UpdateInvoice)

}
