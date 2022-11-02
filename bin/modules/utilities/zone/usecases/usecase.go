package usecases

import (
	"context"
	"trading_be/bin/modules/utilities/zone/helpers"
	"trading_be/bin/modules/utilities/zone/models"
	res "trading_be/bin/pkg/response"
)

func (s *Services) GetList(c context.Context, p *models.ReqGetList) (result res.SendData, err error) {
	result.Data = s.Repository.FindManyCommon(helpers.IDLength[p.Param.Zone], p.Param.ID)

	return result, nil
}
