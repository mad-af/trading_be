package middleware

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("trading_be")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("trading_be")) == 1 {
			return true, nil
		}
		return false, nil
	})
}
