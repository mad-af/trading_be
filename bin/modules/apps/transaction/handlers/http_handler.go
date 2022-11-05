package handlers

import (
	"github.com/labstack/echo/v4"

	m "trading_be/bin/middleware"
	"trading_be/bin/modules/apps/transaction/models"
	"trading_be/bin/modules/apps/transaction/repositories"
	"trading_be/bin/modules/apps/transaction/usecases"
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
	g.POST("/v1/transaction", Create, m.VerifyBearerToken())
	g.PUT("/v1/transaction/:id", Update, m.VerifyBearerToken())
	g.GET("/v1/transaction", GetList, m.VerifyBearerToken())
	g.GET("/v1/transaction/:id", GetDetail, m.VerifyBearerToken())
}

func Create(c echo.Context) error {
	var req = new(models.ReqCreate)
	if err := utils.BindValidate(c, req); err != nil {
		return err
	}
	req.Options = c.Get("opts").(m.JwtClaim)

	var data, err = u.Create(c.Request().Context(), req)
	if err != nil {
		return err
	}
	data.Message = "Create transaction"
	return res.Reply(&data, c)
}

func Update(c echo.Context) error {
	var req = new(models.ReqUpdate)
	if err := utils.BindValidate(c, req); err != nil {
		return err
	}
	req.Options = c.Get("opts").(m.JwtClaim)

	var data, err = u.Update(c.Request().Context(), req)
	if err != nil {
		return err
	}
	data.Message = "Update transaction"
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
	data.Message = "Get List transaction"
	return res.Reply(&data, c)
}

func GetDetail(c echo.Context) error {
	var req = new(models.ReqGetDetail)
	if err := utils.BindValidate(c, req); err != nil {
		return err
	}

	var data, err = u.GetDetail(c.Request().Context(), req)
	if err != nil {
		return err
	}
	data.Message = "Get Detail"
	return res.Reply(&data, c)
}


