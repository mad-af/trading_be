package repositories

import (
	"trading_be/bin/modules/apps/transaction/models"

	"gorm.io/gorm"
)

type (
	Options struct {
		DB *gorm.DB
	}

	Interface interface {
		CreateUser(*models.Users) <-chan res
		Find(*Payload) <-chan res
		Count(*Payload) <-chan res
	}
)
