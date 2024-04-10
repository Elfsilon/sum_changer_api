package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Run() {
	e := echo.New()

	e.POST("/sum", handler)

	e.Logger.Fatal(e.Start(":8080"))
}
