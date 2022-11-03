package models

type Grades struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name" gorm:"unique;not null"`
	Price float64 `json:"price" gorm:"not null"`
}
