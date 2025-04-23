package repository

import (
	"ecommerce-api/internal/domain"
)

var accounts = map[int]*domain.Account{}  

type accountRepository struct{}

func NewAccountRepository() domain.AccountRepository {
	return &accountRepository{}
}

func (r *accountRepository) GetByUserID(userID int) (*domain.Account, error) {
	account, exists := accounts[userID]
	if !exists {
		account = &domain.Account{UserID: userID, Balance: 0}
		accounts[userID] = account
	}
	return account, nil
}

func (r *accountRepository) Update(account *domain.Account) error {
	accounts[account.UserID] = account
	return nil
}
