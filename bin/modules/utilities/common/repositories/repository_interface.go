package repositories

import (
	"gorm.io/gorm"
)

type (
	Options struct {
		DB *gorm.DB
	}

	Interface interface {
		CreateMap(*Payload) <-chan res
		FindMap(*Payload) <-chan res
	}
)
