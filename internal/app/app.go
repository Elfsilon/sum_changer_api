package app

import (
	"sum_changer_api/internal/app/controllers"
	"sum_changer_api/internal/app/middleware"
	"sum_changer_api/internal/app/services"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Run() {
	logger := logrus.New()

	e := echo.New()

	logger.Info("Init dependencies")

	aser := services.NewAccount()
	actr := controllers.NewAccount(aser)

	logger.Info("Init router")
	e.POST("/sum", actr.Handle, middleware.RoleValidator)

	logger.Info("Starting up server")
	e.Logger.Fatal(e.Start(":8080"))

	logger.Info("Server shut down")
}
