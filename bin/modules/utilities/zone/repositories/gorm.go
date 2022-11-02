package repositories

import (
	"fmt"
	"trading_be/bin/modules/utilities/zone/models"
)

func (opt *Options) FindManyCommon(zone int, id string) (result []models.Zone) {
	var like string
	if id != "" {
		like = fmt.Sprintf("id ILIKE '%s%s'", id, "%")
	}
	opt.DB.Table("zones").Where(`length(id) = ?`, zone).Where(like).Find(&result)

	return result
}
