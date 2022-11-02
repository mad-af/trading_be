package usecases

import (
	"context"
	"trading_be/bin/modules/utilities/zone/models"
	res "trading_be/bin/pkg/response"

	"trading_be/bin/modules/utilities/zone/repositories"
)

type (
	Services struct {
		Repository repositories.Interface
	}

	Interface interface {
		GetList(context.Context, *models.ReqGetList) (res.SendData, error)
	}
)
