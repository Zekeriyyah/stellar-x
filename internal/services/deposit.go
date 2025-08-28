// internal/services/deposit_service.go
package services

import (
	"errors"
	"time"

	"github.com/Zekeriyyah/stellar-x/internal/models"
	"github.com/Zekeriyyah/stellar-x/internal/repositories"
)

// Supported stablecoins
var supportedCurrencies = map[string]bool{
	"cNGN": true,
	"cXAF": true,
	"USDx": true,
	"EURx": true,
}

type DepositService struct {
	BalanceRepo *repositories.BalanceRepository
	TransactionRepo *repositories.TransactionRepository
}

func NewDepositService(repo *repositories.BalanceRepository, tx *repositories.TransactionRepository) *DepositService {
	return &DepositService{
		BalanceRepo: repo,
		TransactionRepo: tx,
	}
}

// Deposit simulates a deposit into a wallet
func (s *DepositService) Deposit(walletID uint, currency string, amount float64) (*models.Transaction, error) {
	// Validate input
	if amount <= 0 {
		return nil, errors.New("amount must be positive")
	}

	if !supportedCurrencies[currency] {
		return nil, errors.New("unsupported currency: " + currency)
	}

	// Find Balance
	balance, err := s.BalanceRepo.FindByWalletIDAndCurrency(walletID, currency)
	if err != nil {
		return nil, errors.New("balance not found")
	}

	// update amount in balance
	balance.Amount += amount
	
	// Create transaction record
	transaction := &models.Transaction{
		TxType:       "deposit",
		SenderWalletID: &walletID,
		FromCurrency: currency,
		Amount:       amount,
		CreatedAt:    time.Now(),
	}

	// Save transaction
	if err := s.TransactionRepo.CreateTransaction(transaction); err != nil {
		return nil, errors.New("failed to save transaction")
	}

	return transaction, nil
}