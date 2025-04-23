package repository

import (
	"ecommerce-api/internal/domain"
	"sync"
)

type MockAccountRepository struct {
    accounts map[int]*domain.Account
    mu       sync.Mutex
}

func NewMockAccountRepository() *MockAccountRepository {
    return &MockAccountRepository{
        accounts: make(map[int]*domain.Account),
    }
}

func (r *MockAccountRepository) Create(account *domain.Account) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.accounts[account.UserID] = account
    return nil
}

func (r *MockAccountRepository) GetByUserID(userID int) (*domain.Account, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    account, exists := r.accounts[userID]
    if !exists {
        return nil, nil 
    }
    return account, nil
}

func (r *MockAccountRepository) Update(account *domain.Account) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.accounts[account.UserID] = account
    return nil
}