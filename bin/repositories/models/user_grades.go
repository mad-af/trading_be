package models

import (
	"github.com/google/uuid"
)

type UserGrades struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID `json:"user_id" gorm:"not null"`
	GradeID   int       `json:"grade_id" gorm:"not null"`
	User      Users
	Grade     Grades
}
