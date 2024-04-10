package services

import (
	"sum_changer_api/internal/app/models"
	"sum_changer_api/internal/app/utils"
)

type MockAccount struct {
	sum float32
}

func NewMockAccount() *MockAccount {
	return &MockAccount{
		sum: 1000.0,
	}
}

func (a *MockAccount) Get() (float32, error) {
	return a.sum, nil
}

func (a *MockAccount) HandleOperation(role models.UserRole, sum float32) error {
	if role.IsAdmin() {
		return a.withdraw(sum)
	}

	return a.topUp(sum)
}

func (a *MockAccount) topUp(sum float32) error {
	a.sum += sum
	return nil
}

func (a *MockAccount) withdraw(sum float32) error {
	acc, err := a.Get()
	if err != nil {
		return err
	}

	if utils.Compare(sum, acc) > 0 {
		return ErrInsufficientFunds
	}

	a.sum -= sum
	return nil
}
