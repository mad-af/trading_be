package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HTTPError struct {
	Err     bool   `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewHTTPErrorHandler(err error, c echo.Context) {
	report, ok := err.(*HTTPError)
	if !ok {
		if temp, ok := err.(*echo.HTTPError); ok {
			report = ReplyError(temp.Message.(string), temp.Code)
		} else {
			report = ReplyError(err.Error(), http.StatusInternalServerError)
		}
	}

	c.Logger().Error(report)
	c.JSON(report.Code, report)
}

func ReplyError(message string, code int) *HTTPError {
	err := &HTTPError{Err: true, Code: code, Message: http.StatusText(code)}
	if len(message) > 0 {
		err.Message = message
	}
	return err
}

func (e *HTTPError) Error() string {
	return e.Message
}
