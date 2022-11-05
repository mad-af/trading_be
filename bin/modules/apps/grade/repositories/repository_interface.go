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
		UpdateUserGrade(data *models.UserGrades) <-chan res
		Find(p *Payload) <-chan res
		Count(p *Payload) <-chan res
	}
)
