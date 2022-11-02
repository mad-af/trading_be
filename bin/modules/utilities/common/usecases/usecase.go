package usecases

import (
	"context"
	"encoding/json"
	"trading_be/bin/modules/utilities/common/models"
	rep "trading_be/bin/modules/utilities/common/repositories"
	r "trading_be/bin/pkg/response"
)

func (s *Services) Create(c context.Context, p *models.ReqCreate) (result r.SendData, err error) {
	var res = []map[string]interface{}{}
	result.Data = &res

	var data []map[string]interface{}
	if d, ok := p.Data.(map[string]interface{}); ok {
		data = append(data, d)
	} else {
		var bd, _ = json.Marshal(p.Data)
		json.Unmarshal(bd, &data)
	}

	for _, e := range data {
		var common = <-s.Repository.CreateMap(&rep.Payload{
			Table:    p.Param.Table,
			Document: e,
		})
		if common.Error == nil {
			res = append(res, e)
		}
	}

	return result, nil
}

func (s *Services) GetList(c context.Context, p *models.ReqGetList) (result r.SendData, err error) {
	var res = []map[string]interface{}{}
	result.Data = &res

	var common = <-s.Repository.FindMap(&rep.Payload{
		Table:  p.Param.Table,
		Select: "*",
	})
	if common.Error != nil || common.Row == 0 {
		return result, nil
	}

	res = common.Data.([]map[string]interface{})
	return result, nil
}
