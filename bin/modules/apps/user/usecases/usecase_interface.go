package usecases

import (
	"context"
	"trading_be/bin/modules/apps/user/models"
	res "trading_be/bin/pkg/response"

	"trading_be/bin/modules/apps/user/repositories"
)

type (
	Services struct {
		Repository repositories.Interface
	}

	Interface interface {
		// COMMAND
		Create(context.Context, *models.ReqCreate) (res.SendData, error)
		Login(context.Context, *models.ReqLogin) (res.SendData, error)

		// QUERY
		GetList(context.Context, *models.ReqGetList) (res.SendData, error)
		GetDetail(context.Context, *models.ReqGetDetail) (res.SendData, error)
	}
)
