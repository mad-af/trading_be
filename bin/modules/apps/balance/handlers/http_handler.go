package handlers

import (
	"github.com/labstack/echo/v4"

	m "trading_be/bin/middleware"
	"trading_be/bin/modules/apps/balance/models"
	"trading_be/bin/modules/apps/balance/repositories"
	"trading_be/bin/modules/apps/balance/usecases"
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
	g.POST("/v1/balance/topup", TopUp, m.VerifyBearerToken())
	g.POST("/v1/balance/payment", Payment, m.VerifyBearerToken())
}

func TopUp(c echo.Context) error {
	var req = new(models.ReqTopUp)
	if err := utils.BindValidate(c, req); err != nil {
		return err
	}
	req.Options = c.Get("opts").(m.JwtClaim)

	var data, err = u.TopUp(c.Request().Context(), req)
	if err != nil {
		return err
	}
	data.Message = "Balance top up"
	return res.Reply(&data, c)
}

func Payment(c echo.Context) error {
	var req = new(models.ReqPayment)
	if err := utils.BindValidate(c, req); err != nil {
		return err
	}
	req.Options = c.Get("opts").(m.JwtClaim)

	var data, err = u.Payment(c.Request().Context(), req)
	if err != nil {
		return err
	}
	data.Message = "Balance payment"
	return res.Reply(&data, c)
}