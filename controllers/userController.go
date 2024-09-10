package controllers

import (
	"github.com/AyoOluwa-Israel/invoice-api/db"
	"github.com/AyoOluwa-Israel/invoice-api/models"
	"github.com/AyoOluwa-Israel/invoice-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/asaskevich/govalidator.v9"
)

type response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type RegisterPayload struct {
	Status  int    `json:"status"`
	Message string `json:"message"`

	User models.User `json:"user"`
}

func GetUser(c *fiber.Ctx) error {
	// 1. Get User ID from the request (assuming it's in the URL parameters)
	userID := c.Params("user_id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  fiber.StatusBadRequest,
			Message: "Missing user ID",
		})
	}

	// 2. Retrieve the user from the database
	var user models.User
	if err := db.Database.Db.Preload("PaymentInformation").First(&user, "id = ?", userID).Error; err != nil {

		return c.Status(fiber.StatusNotFound).JSON(response{
			Status:  fiber.StatusNotFound,
			Message: "User not found",
		})

	}

	// 3. Prepare the response data
	data := map[string]interface{}{
		"user": user,
	}

	// 4. Return the user data in the response
	return c.Status(fiber.StatusOK).JSON(response{
		Status:  fiber.StatusOK,
		Message: "User retrieved successfully",
		Data:    data,
	})
}

func GetAllUsers(c *fiber.Ctx) error {
	// 1. Retrieve all users from the database
	var users []models.User
	if err := db.Database.Db.Find(&users).Error; err != nil {
		// Log the actual error for debugging

		return c.Status(fiber.StatusInternalServerError).JSON(response{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to retrieve users",
		})
	}

	// 2. Prepare the response data
	data := map[string]interface{}{
		"users": users,
	}

	// 3. Return the users data in the response
	return c.Status(fiber.StatusOK).JSON(response{
		Status:  fiber.StatusOK,
		Message: "Users retrieved successfully",
		Data:    data,
	})
}

func RegisterUser(c *fiber.Ctx) error {

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response{
			Status:  fiber.StatusUnprocessableEntity,
			Message: "Error parsing data",
		})
	}

	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  fiber.StatusBadRequest,
			Message: "Please provide all fields",
		})
	}

	if !govalidator.IsEmail(user.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  fiber.StatusBadRequest,
			Message: "Email format is wrong!",
		})
	}

	user.Email = utils.ConvertEmail(user.Email)

	exists := db.Database.Db.Where("email = ?", user.Email).First(&user)

	if exists.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  fiber.StatusBadRequest,
			Message: "Email already exists",
		})
	} else {

		user.Id = uuid.New()
		db.Database.Db.Create(&user)
		data := map[string]interface{}{
			"user": user,
		}

		res := response{
			Status:  fiber.StatusCreated,
			Message: "User Created Successfully",
			Data:    data,
		}

		return c.Status(fiber.StatusCreated).JSON(res)
	}

}
