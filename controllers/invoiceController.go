package controllers

import (
	"time"

	"github.com/AyoOluwa-Israel/invoice-api/db"
	"github.com/AyoOluwa-Israel/invoice-api/interfaces"
	"github.com/AyoOluwa-Israel/invoice-api/models"
	"github.com/AyoOluwa-Israel/invoice-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// ShowAccount godoc
//
// @Summary Get invoice by ID
// @Description Retrieve an invoice using its unique ID.
// @Tags Invoice
// @Accept json
// @Produce json
// @Param X-User-Id header string true "User ID"
// @Param id path string true "Invoice ID"
// @Success 200 {object} models.Invoice "Successfully retrieved the invoice"
// @Failure 400 {object} response "Invalid ID format"
// @Failure 404 {object} response "Invoice not found"
// @Failure 500 {object} response "Internal server error"
// @Router /v1/api/invoice/{invoice_id} [get]
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

// ShowAccount godoc
//
// @Summary Get all invoice
// @Description Retrieve all invoice.
// @Tags Invoice
// @Accept json
// @Produce json
// @Param X-User-Id header string true "User ID"
// @Success 200 {object} models.Invoice "Successfully retrieved the invoices"
// @Failure 400 {object} response "Invalid ID format"
// @Failure 404 {object} response "Invoice not found"
// @Failure 500 {object} response "Internal server error"
// @Router /v1/api/invoice [get]
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

// CreateInvoice godoc
//
// @Summary Create an invoice
// @Description Create a new invoice by providing the necessary details.
// @Tags Invoice
// @Accept json
// @Produce json
// @Param X-User-Id header string true "User ID"
// @Param invoice body models.Invoice true "Invoice data"
// @Success 201 {object} models.Invoice "Successfully created the invoice"
// @Failure 400 {object} response "Invalid request data"
// @Failure 500 {object} response "Internal server error"
// @Router /v1/api/invoice [post]
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


// UpdateInvoice godoc
//
// @Summary Update an invoice
// @Description Updates the details of an existing invoice for a specific user.
// @Tags Invoice
// @Accept json
// @Produce json
// @Param invoice_id path string true "Invoice ID"
// @Param X-User-Id header string true "User ID"
// @Param invoice body interfaces.IUpdateInvoice true "Updated invoice data"
// @Success 200 {object} response "Invoice updated successfully"
// @Failure 400 {object} response "Invalid request body"
// @Failure 404 {object} response "Invoice not found"
// @Failure 500 {object} response "Database error or server issue"
// @Router  /v1/api/invoice/{invoice_id} [put]
func UpdateInvoice(c *fiber.Ctx) error {
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

	if err := db.Database.Db.Where("user_id = ? AND invoice_id = ?", userId, invoiceId).First(&invoice).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response{
			Status:  fiber.StatusNotFound,
			Message: "Invoice not found",
		})
	}

	invoice.UpdatedAt = time.Now()

	var updateData interfaces.IUpdateInvoice

	// Parse the request body into the updateData struct
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response{
			Status:  fiber.StatusNotFound,
			Message: "Invalid fields",
		})
	}

	if err := db.Database.Db.Model(&invoice).Updates(updateData).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to update invoice",
		})
	}

	data := map[string]interface{}{
		"invoice": invoice,
	}

	res := response{
		Status:  fiber.StatusCreated,
		Message: "Invoice updated Successfully",
		Data:    data,
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
