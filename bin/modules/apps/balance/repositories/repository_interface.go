package repositories

import (
	"trading_be/bin/modules/apps/balance/models"

	"gorm.io/gorm"
)

type (
	Options struct {
		DB *gorm.DB
	}

	Interface interface {
		UpdateBalance(data *models.Balances) <-chan res
		Find(p *Payload) <-chan res
	}
)
