package models

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	DRAFT           Status = "DRAFT"
	OVERDUE         Status = "OVERDUE"
	PAID            Status = "PAID"
	PENDING_PAYMENT Status = "PENDING_PAYMENT"
	CANCELLED       Status = "CANCELLED"
)

type CurrencyType string

const (
	USD CurrencyType = "USD"
	GBP CurrencyType = "GBP"
	EUR CurrencyType = "EUR"
	NGN CurrencyType = "NGN"
)

type InvoiceActivity struct {
	Action     Status    `json:"action"`
	ActionDate time.Time `json:"action_date"`
}

type Items struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
}

type Reminder string

const (
	TwoWeeks  Reminder = "14 days before due date"
	AWeek     Reminder = "7 days before due date"
	ThreeDays Reminder = "3 days before due date"
	ADay      Reminder = "A day before due date"
	DueDate   Reminder = "Due date"
)

type Invoice struct {
	UserID          uuid.UUID    `json:"user_id" gorm:"not null"`
	InvoiceID       uuid.UUID    `json:"invoice_id" gorm:"type:uuid;primaryKey;not null"`
	InvoiceNumber   string        `json:"invoice_number"`
	Description     string       `json:"description"`
	Status          Status       `json:"status" gorm:"default:'DRAFT'"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	Amount          float64      `json:"amount"`
	DueDate         time.Time    `json:"due_date"`
	BillingCurrency CurrencyType `json:"billing_currency" gorm:"default:'USD'"`
	// Items              []Items           `json:"items" gorm:"default:[]"`
	DiscountPercentage float64 `json:"discount_percentage"`
	Note               string  `json:"note"`
	// Reminders          []Reminder        `json:"reminders" gorm:"default:[]"`
	// InvoiceActivity    []InvoiceActivity `json:"invoice_activity" gorm:"default:[]"`
}
