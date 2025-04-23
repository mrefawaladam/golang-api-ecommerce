package domain

import "sync"

type Account struct {
	ID      int
	UserID  int
	Balance float64
	Mu      sync.RWMutex
}

type AccountRepository interface {
	GetByUserID(userID int) (*Account, error)
	Update(account *Account) error
}

type AccountUsecase interface {
	Deposit(userID int, amount float64) error
	Withdraw(userID int, amount float64) error
	GetBalance(userID int) (float64, error)
}
