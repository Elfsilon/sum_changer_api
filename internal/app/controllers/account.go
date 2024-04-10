package controllers

import (
	"errors"
	"net/http"
	"sum_changer_api/internal/app/models"
	"sum_changer_api/internal/app/services"
	"sum_changer_api/internal/app/utils"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Account struct {
	ser services.AccountService
	log *logrus.Logger
}

func NewAccount(ser services.AccountService, log *logrus.Logger) *Account {
	return &Account{ser, log}
}

func (a *Account) Get(c echo.Context) error {
	acc, err := a.ser.Get()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models.Account{Sum: acc})
}

func (a *Account) Handle(c echo.Context) error {
	role := c.Get(utils.KeyUserRole).(models.UserRole)

	var payload models.Account
	if err := utils.DecodeBody(c, &payload); err != nil {
		return c.String(http.StatusBadRequest, ErrInvalidBody(err).Error())
	}

	if err := c.Bind(&payload); err != nil {
		return c.String(http.StatusBadRequest, ErrInvalidBody(err).Error())
	}

	if payload.Sum <= 0 {
		return c.String(http.StatusBadRequest, ErrNegativeSum.Error())
	}

	if err := a.ser.HandleOperation(role, payload.Sum); err != nil {
		if errors.Is(err, services.ErrInsufficientFunds) {
			return c.String(http.StatusBadRequest, ErrInsufficientFunds.Error())
		}

		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
