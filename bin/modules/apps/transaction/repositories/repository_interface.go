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
		CreateTransaction(*models.Transactions, string) <-chan res
		UpdateStatusTransaction(*models.Transactions, string) <-chan res
		Find(*Payload) <-chan res
		Count(*Payload) <-chan res
	}
)
