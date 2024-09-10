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


// GetUser godoc
//
// @Summary Get a User by ID
// @Description Retrieve a user using its unique ID.
// @Tags User
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 201 {object} models.User "Successfully retrieved the user"
// @Failure 400 {object} response "Invalid request data"
// @Failure 500 {object} response "Internal server error"
// @Router /v1/api/user/{user_id} [get]
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


// GetAllUsers godoc
//
// @Summary Get all Users
// @Description Retrieve a user using its unique ID.
// @Tags User
// @Accept json
// @Produce json
// @Success 201 {object} []models.User "Successfully retrieved all users"
// @Failure 400 {object} response "Invalid request data"
// @Failure 500 {object} response "Internal server error"
// @Router /v1/api/users [get]
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


// RegisterUser godoc
//
// @Summary Create a User
// @Description Create a new user by providing the necessary details.
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.UserPayload true "User data"
// @Success 201 {object} models.User "Successfully created the invoice"
// @Failure 400 {object} response "Invalid request data"
// @Failure 500 {object} response "Internal server error"
// @Router /v1/api/user [post]
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
