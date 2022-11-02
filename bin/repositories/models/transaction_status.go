package models

import (
	"time"

	"github.com/google/uuid"
)

type TransactionStatus struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TransactionID uuid.UUID `json:"transaction_id" gorm:"not null"`
	Status        string    `json:"status" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"not null"`
	Transaction   Transactions
}
