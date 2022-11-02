package utils

import (
	"net/http"
	"trading_be/bin/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ValidationUtil struct {
	validator *validator.Validate
}

func NewValidator() echo.Validator {
	return &ValidationUtil{validator: validator.New()}
}

func (v *ValidationUtil) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func BindValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	if err := c.Validate(i); err != nil {
		return response.ReplyError(err.(validator.ValidationErrors).Error(), http.StatusBadRequest)
	}

	return nil
}
