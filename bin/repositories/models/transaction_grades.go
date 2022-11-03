package models

import (
	"time"

	"github.com/google/uuid"
)

type TransactionUserGrades struct {
	UserGradeID   uuid.UUID `json:"user_grade_id" gorm:"not null"`
	TransactionID uuid.UUID `json:"transaction_id" gorm:"not null"`
	GradeID       int       `json:"grade_id" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"not null"`
	Transaction   Transactions
	UserGrade     UserGrades
	Grade         Grades
}
