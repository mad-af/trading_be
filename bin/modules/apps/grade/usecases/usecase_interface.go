package usecases

import (
	"context"
	"trading_be/bin/modules/apps/grade/models"
	res "trading_be/bin/pkg/response"

	"trading_be/bin/modules/apps/grade/repositories"
)

type (
	Services struct {
		Repository repositories.Interface
	}

	Interface interface {
		// COMMAND
		Transaction(context.Context, *models.ReqTransaction) (res.SendData, error)
		Upgrade(context.Context, *models.ReqUpgrade) (res.SendData, error)
	}
)
