package controllers

import (
	"github.com/AyoOluwa-Israel/invoice-api/db"
	"github.com/AyoOluwa-Israel/invoice-api/models"
	"github.com/AyoOluwa-Israel/invoice-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetInvoiceByID(c *fiber.Ctx) error {
	userId, err := utils.GetUserIDFromHeader(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	invoiceId := c.Params("invoice_id")

	if invoiceId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  fiber.StatusBadRequest,
			Message: "Missing invoice ID",
		})
	}

	var invoice models.Invoice

	if err := db.Database.Db.Where("user_id = ?", userId).First(&invoice, "invoice_id = ?", invoiceId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response{
			Status:  fiber.StatusNotFound,
			Message: "Invoice not found",
		})
	}

	data := map[string]interface{}{
		"invoice": invoice,
	}

	return c.Status(fiber.StatusOK).JSON(response{
		Status:  fiber.StatusOK,
		Message: "Invoice retrieved successfully",
		Data:    data,
	})
}

func GetAllInvoice(c *fiber.Ctx) error {
	userId, err := utils.GetUserIDFromHeader(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var invoice []models.Invoice

	if err := db.Database.Db.Where("user_id = ?", userId).Find(&invoice).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response{
			Status:  fiber.StatusNotFound,
			Message: "Invoice not found",
		})
	}

	data := map[string]interface{}{
		"invoice": invoice,
	}

	return c.Status(fiber.StatusOK).JSON(response{
		Status:  fiber.StatusOK,
		Message: "Invoice retrieved successfully",
		Data:    data,
	})
}

func CreateInvoice(c *fiber.Ctx) error {
	userId, err := utils.GetUserIDFromHeader(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var invoice models.Invoice

	if err := c.BodyParser(&invoice); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response{
			Status:  fiber.StatusUnprocessableEntity,
			Message: "Error parsing data",
		})
	}

	invoice.UserID = userId
	invoice.InvoiceID = uuid.New()
	invoice.InvoiceNumber = utils.AttachRandomNumber("TXID")

	db.Database.Db.Create(&invoice)

	data := map[string]interface{}{
		"invoice": invoice,
	}

	res := response{
		Status:  fiber.StatusCreated,
		Message: "Invoice Created Successfully",
		Data:    data,
	}

	return c.Status(fiber.StatusCreated).JSON(res)

}
