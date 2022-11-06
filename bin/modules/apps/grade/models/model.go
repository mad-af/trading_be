package models

import (
	"time"
	m "trading_be/bin/middleware"
)

// REQUEST
type (
	ReqTransaction struct {
		BankID  int        `json:"bank_id"`
		GradeID int        `json:"grade_id"`
		Options m.JwtClaim `json:"opts"`
	}

	ReqUpgrade struct {
		TransactionID string     `json:"transaction_id"`
		Options       m.JwtClaim `json:"opts"`
	}
)

// REPOSITORY
type (
	TransactionUserGrades struct {
		UserGradeID   string    `json:"user_grade_id"`
		TransactionID string    `json:"transaction_id"`
		GradeID       int       `json:"grade_id"`
		CreatedAt     time.Time `json:"created_at"`
	}

	UserGrades struct {
		ID      string `json:"id"`
		UserID  string `json:"user_id,omitempty"`
		GradeID int    `json:"grade_id"`
	}

	Grades struct {
		ID    uint    `json:"id" gorm:"primaryKey"`
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
)

// COMMON
type ()
