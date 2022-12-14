package main

import (
	"net/http"

	// UTILITIES
	common "trading_be/bin/modules/utilities/common/handlers"
	// zone "trading_be/bin/modules/utilities/zone/handlers"

	// APPS
	user "trading_be/bin/modules/apps/user/handlers"
	transaction "trading_be/bin/modules/apps/transaction/handlers"
	grade "trading_be/bin/modules/apps/grade/handlers"
	balance "trading_be/bin/modules/apps/balance/handlers"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
}

func (s *Server) Routes() {
	e := s.Echo
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
 
	grupUtilities := e.Group("/utilities")
	common.Init(grupUtilities)
	// zone.Init(grupUtilities)

	grupApps := e.Group("/apps")
	user.Init(grupApps)
	transaction.Init(grupApps)
	grade.Init(grupApps)
	balance.Init(grupApps)
}
