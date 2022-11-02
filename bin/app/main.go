package main

import (
	r "trading_be/bin/pkg/response"
	u "trading_be/bin/pkg/utils"
	cf "trading_be/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var s = &Server{echo.New()}
	e := s.Echo

	e.Use(middleware.CORS())

	e.Validator = u.NewValidator()
	e.HTTPErrorHandler = r.NewHTTPErrorHandler

	s.Routes()

	e.Logger.Fatal(e.Start(":" + cf.Env.HttpPort))
}
