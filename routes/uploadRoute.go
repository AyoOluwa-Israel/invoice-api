package routes

import (
	"github.com/AyoOluwa-Israel/invoice-api/config"
	"github.com/AyoOluwa-Israel/invoice-api/utils"
	"github.com/gofiber/fiber/v2"
)

func UploadRoutes(app fiber.Router, config config.Config) {

	uploadHandler := utils.NewUploadHandler(config)

	app.Post("/upload", uploadHandler.UploadToCloudinary)
}
