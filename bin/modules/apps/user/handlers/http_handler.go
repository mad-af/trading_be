package handlers

import (
	"github.com/labstack/echo/v4"

	m "trading_be/bin/middleware"
	"trading_be/bin/modules/apps/user/models"
	"trading_be/bin/modules/apps/user/repositories"
	"trading_be/bin/modules/apps/user/usecases"
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
	g.POST("/v1/user", Create, m.BasicAuth())
	g.POST("/v1/user/login", Login, m.BasicAuth())
	g.PUT("/v1/user/:id", Update, m.VerifyBearerToken())
	g.GET("/v1/user", GetList, m.VerifyBearerToken())
	g.GET("/v1/user/:id", GetDetail, m.VerifyBearerToken())
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
	data.Message = "Create User"
	return res.Reply(&data, c)
}

func Login(c echo.Context) error {
	var req = new(models.ReqLogin)
	if err := utils.BindValidate(c, req); err != nil {
		return err
	}

	var data, err = u.Login(c.Request().Context(), req)
	if err != nil {
		return err
	}
	data.Message = "Login"
	return res.Reply(&data, c)
}

func Update(c echo.Context) error {
	var req = new(models.ReqUpdate)
	if err := utils.BindValidate(c, req); err != nil {
		return err
	}

	var data, err = u.Update(c.Request().Context(), req)
	if err != nil {
		return err
	}
	data.Message = "Update User"
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
	data.Message = "Get List User"
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


