package handlers

import (
	"github.com/labstack/echo/v4"

	m "trading_be/bin/middleware"
	"trading_be/bin/modules/apps/grade/models"
	"trading_be/bin/modules/apps/grade/repositories"
	"trading_be/bin/modules/apps/grade/usecases"
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
	g.POST("/v1/grade/upgrade", Upgrade, m.VerifyBearerToken())
}

func Upgrade(c echo.Context) error {
	var req = new(models.ReqUpgrade)
	if err := utils.BindValidate(c, req); err != nil {
		return err
	}
	req.Options = c.Get("opts").(m.JwtClaim)

	var data, err = u.Upgrade(c.Request().Context(), req)
	if err != nil {
		return err
	}
	data.Message = "Upgrade"
	return res.Reply(&data, c)
}
