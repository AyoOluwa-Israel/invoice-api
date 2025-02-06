package utils

import (
	"github.com/AyoOluwa-Israel/invoice-api/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
)

type response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type UploadHandler struct {
	Config config.Config
}

func NewUploadHandler(config config.Config) *UploadHandler {
	return &UploadHandler{Config: config}
}

func (h *UploadHandler) UploadToCloudinary(c *fiber.Ctx) error {

	cloudName := h.Config.CloudinaryCloudName
	apiKey := h.Config.CloudinaryApiKey
	apiSecret := h.Config.CloudinarySecretKey

	file, err := c.FormFile("file")

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  fiber.StatusBadRequest,
			Message: "Unable to parse file",
		})
	}

	fileReader, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response{
			Status:  fiber.StatusInternalServerError,
			Message: "Unable to open file",
		})
	}

	defer fileReader.Close()

	// Initialize Cloudinary
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(response{
			Status:  fiber.StatusBadGateway,
			Message: "Unable to open file",
		})
	}
	uploadResult, err := cld.Upload.Upload(c.Context(), fileReader, uploader.UploadParams{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to upload file to Cloudinary",
		})
	}

	data := map[string]interface{}{
		"url": uploadResult.SecureURL,
	}

	return c.JSON(response{
		Status:  fiber.StatusCreated,
		Message: "Image Uploaded Successfully",
		Data:    data,
	})

}
