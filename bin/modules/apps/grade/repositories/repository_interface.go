package repositories

import (
	"trading_be/bin/modules/apps/grade/models"

	"gorm.io/gorm"
)

type (
	Options struct {
		DB *gorm.DB
	}

	Interface interface {
		CreateTransactionUserGrade(*models.TransactionUserGrades) <-chan res
		Find(p *Payload) <-chan res
	}
)
