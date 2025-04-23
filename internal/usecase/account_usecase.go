package usecase

import (
	"ecommerce-api/internal/domain"
	"errors"
)

type accountUsecase struct {
	repo domain.AccountRepository
}

func NewAccountUsecase(repo domain.AccountRepository) domain.AccountUsecase {
	return &accountUsecase{repo: repo}
}

func (u *accountUsecase) Deposit(userID int, amount float64) error {
	account, _ := u.repo.GetByUserID(userID)
	account.Mu.Lock()
	defer account.Mu.Unlock()
	account.Balance += amount
	return u.repo.Update(account)
}

func (u *accountUsecase) Withdraw(userID int, amount float64) error {
	account, _ := u.repo.GetByUserID(userID)
	account.Mu.Lock()
	defer account.Mu.Unlock()
	if account.Balance < amount {
		return errors.New("saldo tidak cukup")
	}
	account.Balance -= amount
	return u.repo.Update(account)
}

func (u *accountUsecase) GetBalance(userID int) (float64, error) {
	account, _ := u.repo.GetByUserID(userID)
	account.Mu.RLock()
	defer account.Mu.RUnlock()
	return account.Balance, nil
}


