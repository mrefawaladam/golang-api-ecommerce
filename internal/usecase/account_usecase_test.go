package usecase_test

import (
	"ecommerce-api/internal/repository"
	"ecommerce-api/internal/usecase"
	"sync"
	"testing"
	"time"
)

func TestConcurrentDepositWithdraw(t *testing.T) {
	repo := repository.NewAccountRepository()
	uc := usecase.NewAccountUsecase(repo)

	userID := 1

	var wg sync.WaitGroup

	// Simulasi 100 Deposit (masing-masing 10)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := uc.Deposit(userID, 10)
			if err != nil {
				t.Error("Deposit error:", err)
			}
		}()
	}

	// Simulasi 50 Withdraw (masing-masing 5)
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := uc.Withdraw(userID, 5)
			if err != nil {
				// Bisa error kalau saldo belum cukup, jadi kita abaikan yang ini
			}
		}()
	}

	wg.Wait()
	time.Sleep(100 * time.Millisecond) // Pastikan semua goroutine selesai

	balance, err := uc.GetBalance(userID)
	if err != nil {
		t.Fatal("Gagal ambil saldo:", err)
	}

	expectedMin := float64(500) // 100 * 10 - 50 * 5 = 1000 - 250 (min)
	if balance < expectedMin {
		t.Errorf("Saldo akhir terlalu kecil: got %v, want at least %v", balance, expectedMin)
	}
}
