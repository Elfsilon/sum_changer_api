package services

import "sum_changer_api/internal/app/models"

type AccountService interface {
	Get() (float64, error)
	HandleOperation(role models.UserRole, sum float64) error
}
