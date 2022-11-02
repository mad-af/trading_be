package repositories

import (
	"trading_be/bin/modules/utilities/zone/models"

	"gorm.io/gorm"
)

type (
	Options struct {
		DB *gorm.DB
	}

	Interface interface {
		FindManyCommon(zone int, id string) []models.Zone
	}
)
