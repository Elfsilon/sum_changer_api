package repos

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type Account struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewAccount(db *sql.DB, log *logrus.Logger) *Account {
	return &Account{db, log}
}

func (a *Account) Get() (float32, error) {
	query := `SELECT sum FROM Account;`

	var sum float32

	err := a.db.QueryRow(query).Scan(&sum)
	if err != nil {
		return 0.0, err
	}

	return sum, nil
}

func (a *Account) TopUp(sum float32) error {
	query := `UPDATE Account SET sum = sum + $1;`

	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	err = tx.QueryRow(query, sum).Err()
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (a *Account) Withdraw(sum float32) error {
	query := `UPDATE Account SET sum = sum - $1;`

	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	err = tx.QueryRow(query, sum).Err()
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
