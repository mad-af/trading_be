package models

import (
	"github.com/google/uuid"
)

type Balances struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID uuid.UUID `json:"user_id" gorm:"not null"`
	Value  int64     `json:"value" gorm:"not null"`
	User   Users
}
