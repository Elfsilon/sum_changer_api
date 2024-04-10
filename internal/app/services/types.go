package services

import "sum_changer_api/internal/app/models"

type AccountService interface {
	Get() (float32, error)
	HandleOperation(role models.UserRole, sum float32) error
}
