package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sum_changer_api/internal/app/controllers"
	"sum_changer_api/internal/app/middleware"
	repos "sum_changer_api/internal/app/repositories"
	"sum_changer_api/internal/app/services"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var (
	dbpass   = MustGetEnv("DB_PASSWORD")
	logLevel = MustGetEnv("LOG_LEVEL")
)

func MustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("env variable %v not set", key)
	}
	return val
}

func setupLogger(logLevel string) *logrus.Logger {
	logger := logrus.New()

	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logger.Warningf("cannot parse passed log level, set default")
		lvl = logrus.InfoLevel
	}
	logger.SetLevel(lvl)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})

	return logger
}

func Run() {
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
