package main

import (
	"log"

	"github.com/AyoOluwa-Israel/invoice-api/config"
	"github.com/AyoOluwa-Israel/invoice-api/db"
	"github.com/AyoOluwa-Israel/invoice-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
			// AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
			AllowOrigins: "*",
			// AllowCredentials: true,
			// AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		}))

	router := app.Group("/v1/api")
	routes.UserRoutes(router)
	routes.PaymentInformationRoutes(router)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Welcome to my Invoice API",
			"status":  fiber.StatusOK,
		})
	})

	log.Fatal(app.Listen(":" + config.ServerPort))

}
