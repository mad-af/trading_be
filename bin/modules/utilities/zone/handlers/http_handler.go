package handlers

import (
	"github.com/labstack/echo/v4"

	// "trading_be/bin/middleware"
	"trading_be/bin/modules/utilities/zone/models"
	"trading_be/bin/modules/utilities/zone/repositories"
	"trading_be/bin/modules/utilities/zone/usecases"
	res "trading_be/bin/pkg/response"
	"trading_be/bin/pkg/utils"
	db "trading_be/bin/repositories"
)

var gorm repositories.Interface = &repositories.Options{
	DB: db.InitPostgre(),
}

var u usecases.Interface = &usecases.Services{
	Repository: gorm,
}

func Init(g *echo.Group) {
	g.GET("/v1/zone/:zone", GetList)
	g.GET("/v1/zone/:zone/:id", GetList)
}

func GetList(c echo.Context) error {
	var req = new(models.ReqGetList)
	if err := utils.BindValidate(c, req); err != nil {
		return err
	}

	var data, err = u.GetList(c.Request().Context(), req)
	if err != nil {
		return err
	}
	return res.Reply(&data, c)
}
