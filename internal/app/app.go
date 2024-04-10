package app

import (
	"database/sql"
	"fmt"
	"os"
	"sum_changer_api/internal/app/controllers"
	"sum_changer_api/internal/app/middleware"
	repos "sum_changer_api/internal/app/repositories"
	"sum_changer_api/internal/app/services"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

func setupLogger(logLevel string) *logrus.Logger {
	logger := logrus.New()

	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logger.Warningf("cannot parse passed log level, set default")
		lvl = logrus.InfoLevel
	}
	logger.SetLevel(lvl)

	return logger
}

func Run() {
	dbpass := os.Getenv("DB_PASSWORD")
	logLevel := os.Getenv("LOG_LEVEL")

	logger := setupLogger(logLevel)

	// Init db
	dbconn := "postgresql://postgres:%v@db:5432/postgres?sslmode=disable"
	dbconn = fmt.Sprintf(dbconn, dbpass)

	db, err := sql.Open("postgres", dbconn)
	if err != nil {
		logger.Fatalf("error while opening database: %s", err)
	}
	defer db.Close()

	logger.Info("Init dependencies")

	arep := repos.NewAccount(db, logger)
	aser := services.NewAccount(arep, logger)
	actr := controllers.NewAccount(aser, logger)

	// Init server & router
	e := echo.New()

	e.GET("/sum", actr.Get)
	e.POST("/sum", actr.Handle, middleware.RoleValidator)

	e.Logger.Fatal(e.Start(":8080"))

	logger.Info("Server shut down")
}
