package controllers

import "errors"

var (
	ErrNegativeSum       = errors.New("sum must be greater than 0")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

func ErrInvalidBody(err error) error {
	return errors.Join(errors.New("invalid body: "), err)
}
