package services

import (
	"errors"

	"github.com/Zekeriyyah/stellar-x/internal/models"
	"github.com/Zekeriyyah/stellar-x/internal/repositories"
	"gorm.io/gorm"
)

type BalanceService struct {
	BalanceRepo *repositories.BalanceRepository
}

func NewBalanceService(balRepo *repositories.BalanceRepository) *BalanceService {
	return &BalanceService{
		BalanceRepo: balRepo,
	}
}

// InitializeBalances creates zero balances for all supported stablecoins
func (s *BalanceService) InitializeBalances(walletID uint) error {
	balances := []models.Balance{
		{WalletID: walletID, Currency: "cNGN", Amount: 0},
		{WalletID: walletID, Currency: "cXAF", Amount: 0},
		{WalletID: walletID, Currency: "USDx", Amount: 0},
		{WalletID: walletID, Currency: "EURx", Amount: 0},
	}
	return s.BalanceRepo.InitBalances(balances)
}

// GetBalance returns the current amount for a currency
func (s *BalanceService) GetBalance(walletID uint, currency string) (*models.Balance, error) {
	return s.BalanceRepo.FindByWalletIDAndCurrency(walletID, currency)
}

// Increase adds to a balance
func (s *BalanceService) Increase(walletID uint, currency string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	balance, err := s.GetBalance(walletID, currency)
	if err != nil {
		return errors.New("could not get balance for the wallet")
	}

	balance.Amount += amount
	return s.BalanceRepo.Update(balance)
}

// Decrease subtracts from a balance (fails if insufficient)
func (s *BalanceService) Decrease(walletID uint, currency string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	balance, err := s.GetBalance(walletID, currency)
	if err != nil {
		return errors.New("invalid currency or wallet id")
	}

	if balance.Amount < amount {
		return errors.New("insufficient balance")
	}

	balance.Amount -= amount

	// ✅ Save updated balance to database
	if err := s.BalanceRepo.Update(balance); err != nil {
		return errors.New("failed to update balance")
	}
	return nil
}

// UpdateInTx updates balance within a transaction (for swaps/transfers)
func (s *BalanceService) UpdateInTx(tx *gorm.DB, walletID uint, currency string, delta float64) error {
	balance, err := s.BalanceRepo.FindByWalletIDAndCurrency(walletID, currency)
	if err != nil {
		return err
	}

	newAmount := balance.Amount + delta
	if newAmount < 0 {
		return errors.New("insufficient balance")
	}

	balance.Amount = newAmount
	return s.BalanceRepo.UpdateInTx(tx, balance)
}