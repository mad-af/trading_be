package usecases

import (
	"context"
	"trading_be/bin/modules/apps/balance/models"
	res "trading_be/bin/pkg/response"

	"trading_be/bin/modules/apps/balance/repositories"
)

type (
	Services struct {
		Repository repositories.Interface
	}

	Interface interface {
		// COMMAND
		TopUp(context.Context, *models.ReqTopUp) (res.SendData, error)
	}
)
