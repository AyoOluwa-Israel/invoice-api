package controllers

import (
	"github.com/AyoOluwa-Israel/invoice-api/db"
	"github.com/AyoOluwa-Israel/invoice-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetInvoiceByID(c *fiber.Ctx) error {
	invoiceId := c.Params("invoice_id")

	if invoiceId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  fiber.StatusBadRequest,
			Message: "Missing invoice ID",
		})
	}

	var invoice models.Invoice

	if err := db.Database.Db.First(&invoice, "invoice_id = ?", invoiceId).Error; err != nil {
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

	var invoice models.Invoice

	invoice.InvoiceID = uuid.New()
	db.Database.Db.Create(&invoice)

	data := map[string]interface{}{
		"user": invoice,
	}

	res := response{
		Status:  fiber.StatusCreated,
		Message: "Invoice Created Successfully",
		Data:    data,
	}

	return c.Status(fiber.StatusCreated).JSON(res)

}
