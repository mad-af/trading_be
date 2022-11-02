package usecases

import (
	"context"
	"trading_be/bin/modules/utilities/common/models"
	res "trading_be/bin/pkg/response"

	"trading_be/bin/modules/utilities/common/repositories"
)

type (
	Services struct {
		Repository repositories.Interface
	}

	Interface interface {
		Create(context.Context, *models.ReqCreate) (res.SendData, error)
		GetList(context.Context, *models.ReqGetList) (res.SendData, error)
	}
)
