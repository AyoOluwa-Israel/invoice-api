package controllers

import (
	"log"

	"github.com/AyoOluwa-Israel/invoice-api/db"
	"github.com/AyoOluwa-Israel/invoice-api/models"
	"github.com/AyoOluwa-Israel/invoice-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/asaskevich/govalidator.v9"
)

func PostMessage(c *fiber.Ctx) error {


	var message models.MessageStruct

	if err := c.BodyParser(&message); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response{
			Status:  fiber.StatusUnprocessableEntity,
			Message: "Error parsing data",
		})
	}

	if message.Message == "" || message.Name == "" || message.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  fiber.StatusBadRequest,
			Message: "Please provide all fields",
		})
	}

	if !govalidator.IsEmail(message.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  fiber.StatusBadRequest,
			Message: "Email format is wrong!",
		})
	}

	if err := utils.SendEmail(message); err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}
	// if err != nil {
	// 	return c.Status(fiber.StatusUnprocessableEntity).JSON(response{
	// 		Status:  fiber.StatusUnprocessableEntity,
	// 		Message: "Error sending email",
	// 	})
	// }

	message.Id = uuid.New()
	db.Database.Db.Create(&message)

	res := response{
		Status:  fiber.StatusCreated,
		Message: "Thank you for contacting me!",
	}

	return c.Status(fiber.StatusCreated).JSON(res)

}
