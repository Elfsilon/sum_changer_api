package controllers

import (
	"errors"
	"log"
	"net/http"
	"sum_changer_api/internal/app/models"
	"sum_changer_api/internal/app/services"
	"sum_changer_api/internal/app/utils"

	"github.com/labstack/echo/v4"
)

type Account struct {
	ser services.AccountService
}

func NewAccount(ser services.AccountService) *Account {
	return &Account{ser}
}

func (a *Account) Handle(c echo.Context) error {
	role := c.Get(utils.KeyUserRole).(models.UserRole)

	var payload models.Request
	if err := utils.DecodeBody(c, &payload); err != nil {
		return c.String(http.StatusBadRequest, ErrInvalidBody(err).Error())
	}

	if err := c.Bind(&payload); err != nil {
		return c.String(http.StatusBadRequest, ErrInvalidBody(err).Error())
	}

	log.Printf("req sum is %v", payload.Sum)

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
