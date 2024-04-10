package services

import (
	"errors"
	"sum_changer_api/internal/app/models"
	repo "sum_changer_api/internal/app/repositories"
	"sum_changer_api/internal/app/utils"

	"github.com/sirupsen/logrus"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type Account struct {
	rep repo.AccountRepository
	log *logrus.Logger
}

func NewAccount(rep repo.AccountRepository, log *logrus.Logger) *Account {
	return &Account{rep, log}
}

func (a *Account) Get() (float32, error) {
	return a.rep.Get()
}

func (a *Account) HandleOperation(role models.UserRole, sum float32) error {

	if role.IsAdmin() {
		return a.withdraw(sum)
	}

	return a.topUp(sum)
}

func (a *Account) topUp(sum float32) error {
	a.log.Infof("Top up account by %v", sum)
	return a.rep.TopUp(sum)
}

func (a *Account) withdraw(sum float32) error {
	a.log.Infof("Withdraw %v from account", sum)

	acc, err := a.Get()
	if err != nil {
		return err
	}

	if utils.Compare(sum, acc) > 0 {
		return ErrInsufficientFunds
	}

	return a.rep.Withdraw(sum)
}
