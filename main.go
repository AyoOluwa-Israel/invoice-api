package main

import (
	"log"

	"github.com/AyoOluwa-Israel/invoice-api/config"
	"github.com/AyoOluwa-Israel/invoice-api/db"
	_ "github.com/AyoOluwa-Israel/invoice-api/docs"
	"github.com/AyoOluwa-Israel/invoice-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func init() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	db.NewConnection(&config)
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	app := fiber.New(fiber.Config{
		AppName:       "Welcome to the Invoice Api",
		CaseSensitive: true,
	})

	app.Use(logger.New())

	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
		}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	router := app.Group("/v1/api")
	routes.UserRoutes(router)
	routes.WebsiteRoutes(router)
	routes.PaymentInformationRoutes(router)
	routes.InvoiceRoutes(router)

	// Swagger UI route

	app.Get("/", func(c *fiber.Ctx) error {

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Welcome to my Invoice API",
			"status":  fiber.StatusOK,
		})
	})

	log.Fatal(app.Listen(":" + config.ServerPort))

}
