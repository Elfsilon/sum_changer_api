package middleware

import (
	"net/http"
	"sum_changer_api/internal/app/models"
	"sum_changer_api/internal/app/utils"

	"github.com/labstack/echo/v4"
)

var acceptedRoles = map[string]struct{}{
	"admin":  {},
	"client": {},
}

func RoleValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		roles, ok := c.Request().Header["User-Role"]
		if !ok || len(roles) == 0 {
			return c.String(http.StatusForbidden, "access denied")
		}

		if _, ok := acceptedRoles[roles[0]]; !ok {
			return c.String(http.StatusForbidden, "access denied")
		}

		role := models.UserRole(roles[0])
		c.Set(utils.KeyUserRole, role)

		return next(c)
	}
}
