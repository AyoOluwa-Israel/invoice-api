package models

import (
	"time"

	"github.com/google/uuid"
)

type PaymentInformation struct {
	ID                   uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;"`
	AccountName          string    `json:"account_name"`
	AccountNumber        string    `json:"account_number"`
	AccountRoutingNumber string    `json:"account_routing_number"`
	BankName             string    `json:"bank_name"`
	UserID               uuid.UUID `json:"user_id" gorm:"not null"`
}

type User struct {
	Id                 uuid.UUID            `json:"id" gorm:"primaryKey;type:uuid;"`
	FirstName          string               `json:"first_name"`
	LastName           string               `json:"last_name"`
	Phone              string               `json:"phone"`
	Email              string               `json:"email"`
	ImageUrl           string               `json:"image_url"`
	Verified           bool                 `json:"is_verified" `
	CreatedAt          time.Time            `json:"created_at"`
	UpdatedAt          time.Time            `json:"updated_at"`
	PaymentInformation []PaymentInformation `json:"payment_info" gorm:"foreignKey:UserID;default:[]"`
	Invoice []Invoice `json:"invoice" gorm:"foreignKey:UserID;default:[]"`
}


type UserPayload struct {
	FirstName          string               `json:"first_name"`
	LastName           string               `json:"last_name"`
	Phone              string               `json:"phone"`
	Email              string               `json:"email"`
}
