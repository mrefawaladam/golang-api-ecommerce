package usecase

import (
	"ecommerce-api/internal/domain"
	"errors"
	"sync"
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
func (u *accountUsecase) SimulateConcurrent(userID int, initialBalance, depositAmount, withdrawAmount float64, numGoroutines int) (float64, float64, error) {
	 
	account, err := u.repo.GetByUserID(userID)
	if err != nil || account == nil {
		account = &domain.Account{
			UserID:  userID,
			Balance: 0,
		}
	}
	account.Mu.Lock()
	account.Balance = initialBalance
	account.Mu.Unlock()
	if err := u.repo.Update(account); err != nil {
		return 0, 0, err
	}

 	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			_ = u.Deposit(userID, depositAmount)
		}()
		go func() {
			defer wg.Done()
			_ = u.Withdraw(userID, withdrawAmount)
		}()
	}
	wg.Wait()

	finalBalance, _ := u.GetBalance(userID)
	expected := initialBalance + float64(numGoroutines)*(depositAmount-withdrawAmount)

	return finalBalance, expected, nil
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


