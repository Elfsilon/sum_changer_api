package services

import (
	"errors"
	"log"
	"math/big"
	"sum_changer_api/internal/app/models"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type Account struct{}

func NewAccount() *Account {
	return &Account{}
}

func (a *Account) Get() (float64, error) {
	return 1000, nil
}

func (a *Account) HandleOperation(role models.UserRole, sum float64) error {

	if role.IsAdmin() {
		return a.withdraw(sum)
	}

	return a.topUp(sum)
}

func (a *Account) topUp(sum float64) error {
	log.Println("TOP UP")

	// TODO: Top up

	return nil
}

func (a *Account) withdraw(sum float64) error {
	log.Println("WITHDRAW")

	acc, err := a.Get()
	if err != nil {
		return err
	}

	if big.NewFloat(sum).Cmp(big.NewFloat(acc)) > 0 {
		return ErrInsufficientFunds
	}

	// TODO: Withdraw

	return nil
}
