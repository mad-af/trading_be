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
		Upgrade(context.Context, *models.ReqUpgrade) (res.SendData, error)
	}
)
