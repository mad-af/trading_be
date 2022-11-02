package models

import "github.com/google/uuid"

type Transactions struct {
	ID                uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID            uuid.UUID `json:"user_id" gorm:"not null"`
	BankID            int       `json:"bank_id" gorm:"not null"`
	TransactionTypeID int       `json:"transaction_type_id" gorm:"not null"`
	Value             int64     `json:"value" gorm:"not null"`
	Description       string    `json:"description" gorm:"not null"`
	User              Users
	Bank              Banks
	TransactionType   TransactionTypes
}
