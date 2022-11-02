package handlers

import (
	"github.com/labstack/echo/v4"

	m "trading_be/bin/middleware"
	"trading_be/bin/modules/utilities/common/models"
	"trading_be/bin/modules/utilities/common/repositories"
	"trading_be/bin/modules/utilities/common/usecases"
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
	g.POST("/v1/common/:table", Create, m.BasicAuth())
	g.GET("/v1/common/:table", GetList, m.BasicAuth())
}

func Create(c echo.Context) error {
	var req = new(models.ReqCreate)
	if err := utils.BindValidate(c, req); err != nil {
		return err
	}

	var data, err = u.Create(c.Request().Context(), req)
	if err != nil {
		return err
	}
	data.Message = "Create data table " + req.Param.Table
	return res.Reply(&data, c)
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
	data.Message = "get list common table " + req.Param.Table
	return res.Reply(&data, c)
}
