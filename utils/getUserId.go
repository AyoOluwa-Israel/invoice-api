package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Extracts and validates the UserID from the header
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

	return userID, nil
}