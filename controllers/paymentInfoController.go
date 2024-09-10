package controllers

import (
	"fmt"


	"github.com/AyoOluwa-Israel/invoice-api/db"
	"github.com/AyoOluwa-Israel/invoice-api/models"
	"github.com/AyoOluwa-Israel/invoice-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreatePaymentInfo(c *fiber.Ctx) error {

	userId, err := utils.GetUserIDFromHeader(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Check if the UserID exists in the users table
	var user models.User
	if err := db.Database.Db.First(&user, "id = ?", userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).SendString("User not found")
		}
		// Log any other error
		fmt.Println("Error finding user:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Database error")
	}



	var paymentInfo models.PaymentInformation

	if err := c.BodyParser(&paymentInfo); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response{
			Status:  fiber.StatusUnprocessableEntity,
			Message: "Error parsing data",
		})
	}

	paymentInfo.UserID = userId
	paymentInfo.ID = uuid.New()

	db.Database.Db.Create(&paymentInfo)

	data := map[string]interface{}{
		"payment_info": paymentInfo,
	}

	res := response{
		Status:  fiber.StatusCreated,
		Message: "Payment information Created Successfully",
		Data:    data,
	}

	return c.Status(fiber.StatusCreated).JSON(res)

}

func GetAllPaymentInfo(c *fiber.Ctx) error {

	userId, err := utils.GetUserIDFromHeader(c)


	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var paymentInfo []models.PaymentInformation
	if err := db.Database.Db.Where("user_id = ?", userId).Find(&paymentInfo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to retrieve payment info",
		})
	}

	data := map[string]interface{}{
		"payment_info": paymentInfo,
	}

	return c.Status(fiber.StatusOK).JSON(response{
		Status:  fiber.StatusOK,
		Message: "Payment info retrieved successfully",
		Data:    data,
	})

}
