package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	ReplySend struct {
		Err     bool        `json:"error"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}

	ReplyMetaSend struct {
		Err     bool        `json:"error"`
		Data    interface{} `json:"data"`
		Meta    *Meta       `json:"meta"`
		Message string      `json:"message"`
	}

	SendData struct {
		Code    int
		Data    interface{} `json:"data"`
		Meta    *Meta       `json:"meta"`
		Message string      `json:"message"`
	}

	Meta struct {
		Page      int `json:"page"`
		Quantity  int `json:"quantity"`
		TotalPage int `json:"total_page"`
		TotalData int `json:"total_data"`
	}
)

func Reply(data *SendData, c echo.Context) error {
	var final interface{}

	if data.Meta == nil {
		final = ReplySend{
			Err:     false,
			Data:    data.Data,
			Message: data.Message,
		}
	} else {
		final = ReplyMetaSend{
			Err:     false,
			Data:    data.Data,
			Meta:    data.Meta,
			Message: data.Message,
		}
	}

	return c.JSON(data.Code, final)
}

func HTML(html *string, c echo.Context) error {
	return c.HTML(http.StatusOK, *html)
}
