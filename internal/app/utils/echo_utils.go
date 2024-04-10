package utils

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func DecodeBody(c echo.Context, v any) error {
	return json.NewDecoder(c.Request().Body).Decode(&v)
}
