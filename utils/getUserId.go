package utils

import (
	"github.com/AyoOluwa-Israel/invoice-api/db"
	"github.com/AyoOluwa-Israel/invoice-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)




func GetUserIDFromHeader(c *fiber.Ctx) (uuid.UUID, error) {
	// Extract UserID from the header
	id := c.Get("X-User-Id")

	// Check if the header is missing
	if id == "" {
		return uuid.Nil, fiber.NewError(fiber.StatusBadRequest, "User ID is required")
	}

	// Parse the UserID to UUID
	userID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, fiber.NewError(fiber.StatusBadRequest, "Invalid User ID format")
	}

	// Check if the UserID exists in the users table
	var user models.User
	if err := db.Database.Db.First(&user, "id = ?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return uuid.Nil, fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		// Log any other error
		return uuid.Nil, fiber.NewError(fiber.StatusInternalServerError, "Database error")
	}

	return userID, nil
}