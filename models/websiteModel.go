package models

import (
	"time"

	"github.com/google/uuid"
)

type MessageStruct struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Message   string    `json:"message"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
}
