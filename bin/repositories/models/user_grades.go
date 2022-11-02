package models

import (
	"time"

	"github.com/google/uuid"
)

type UserGrades struct {
	UserID    uuid.UUID `json:"user_id" gorm:"not null"`
	GradeID   int       `json:"grade_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
	User      Users
	Grade     Grades
}
