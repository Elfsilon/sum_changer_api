package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sum_changer_api/internal/app/middleware"
	"sum_changer_api/internal/app/services"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetAccount(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/sum", nil)
	res := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, res)

	aser := services.NewMockAccount()
	ctr := NewAccount(aser, logrus.New())

	if assert.NoError(t, ctr.Get(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "{\"sum\":1000}", strings.Trim(res.Body.String(), "\n"))
	}
}

type handleTestConfig struct {
	payload          string
	role             string
	setRoleHeader    bool
	setRoleValidator bool
}

func setupHandle(cfg handleTestConfig) (*httptest.ResponseRecorder, services.AccountService, error) {
	req := httptest.NewRequest(http.MethodPost, "/sum", strings.NewReader(cfg.payload))
	if cfg.setRoleHeader {
		req.Header.Set("User-Role", cfg.role)
	}
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, res)

	aser := services.NewMockAccount()
	ctr := NewAccount(aser, logrus.New())
	h := ctr.Handle
	if cfg.setRoleValidator {
		h = middleware.RoleValidator(h)
	}

	return res, aser, h(c)
}

func TestHandleClientNegativeSum(t *testing.T) {
	res, _, err := setupHandle(handleTestConfig{
		payload:          `{"sum":-100}`,
		role:             "client",
		setRoleHeader:    true,
		setRoleValidator: true,
	})

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "sum must be greater than 0", res.Body.String())
	}
}

func TestHandleAdminNegativeSum(t *testing.T) {
	res, _, err := setupHandle(handleTestConfig{
		payload:          `{"sum":-100}`,
		role:             "admin",
		setRoleHeader:    true,
		setRoleValidator: true,
	})

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "sum must be greater than 0", res.Body.String())
	}
}

func TestHandleTopUp(t *testing.T) {
	res, ser, err := setupHandle(handleTestConfig{
		payload:          `{"sum":100}`,
		role:             "client",
		setRoleHeader:    true,
		setRoleValidator: true,
	})

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, res.Code)
	}

	acc, err := ser.Get()
	if assert.NoError(t, err) {
		assert.Equal(t, float32(1100), acc)
	}
}

func TestHandleWithdraw(t *testing.T) {
	res, ser, err := setupHandle(handleTestConfig{
		payload:          `{"sum":100}`,
		role:             "admin",
		setRoleHeader:    true,
		setRoleValidator: true,
	})

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, res.Code)
	}

	acc, err := ser.Get()
	if assert.NoError(t, err) {
		assert.Equal(t, float32(900), acc)
	}
}

func TestHandleWithdrawInsufficientFunds(t *testing.T) {
	res, _, err := setupHandle(handleTestConfig{
		payload:          `{"sum":1100}`,
		role:             "admin",
		setRoleHeader:    true,
		setRoleValidator: true,
	})

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusForbidden, res.Code)
		assert.Equal(t, "insufficient funds", res.Body.String())
	}
}

func TestHandleHeaderless(t *testing.T) {
	res, _, err := setupHandle(handleTestConfig{
		payload:          `{"sum":100}`,
		role:             "admin",
		setRoleHeader:    false,
		setRoleValidator: true,
	})

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusForbidden, res.Code)
		assert.Equal(t, "access denied", res.Body.String())
	}
}

func TestHandleAdminWithoutRoleValidator(t *testing.T) {
	assert.Panics(t, func() {
		setupHandle(handleTestConfig{
			payload:          `{"sum":100}`,
			role:             "admin",
			setRoleHeader:    true,
			setRoleValidator: false,
		})
	})
}

func TestHandleClientWithoutRoleValidator(t *testing.T) {
	assert.Panics(t, func() {
		setupHandle(handleTestConfig{
			payload:          `{"sum":100}`,
			role:             "client",
			setRoleHeader:    true,
			setRoleValidator: false,
		})
	})
}
