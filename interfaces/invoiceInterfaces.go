package interfaces

import (
	"time"

	"github.com/AyoOluwa-Israel/invoice-api/models"
)

type IUpdateInvoice struct {
	Description        string              `json:"description"`
	Amount             float64             `json:"amount"`
	DueDate            time.Time           `json:"due_date"`
	Note               string              `json:"note"`
	DiscountPercentage float64             `json:"discount_percentage"`
	BillingCurrency    models.CurrencyType `json:"billing_currency" gorm:"default:'USD'"`
	UpdatedAt       time.Time    `json:"updated_at"`
}
