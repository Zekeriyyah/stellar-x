package services

import (
	"errors"
	"time"

	"github.com/Zekeriyyah/stellar-x/internal/models"
	"github.com/Zekeriyyah/stellar-x/internal/repositories"
	"gorm.io/gorm"
)

type SwapService struct {
	BalanceService   *BalanceService
	TransactionRepo  *repositories.TransactionRepository
	FXService        *FXService
}

func NewSwapService(b *BalanceService,	tx *repositories.TransactionRepository,	fx *FXService) *SwapService {
	return &SwapService{
		BalanceService:  b,
		TransactionRepo: tx,
		FXService:       fx,
	}
}

// Swap executes a cross-currency swap
func (s *SwapService) Swap(walletID uint, fromCurrency, toCurrency string,	amount float64) (*models.Transaction, error) {
	// Validate input
	if amount <= 0 {
		return nil, errors.New("amount must be positive")
	}

	if fromCurrency == toCurrency {
		return nil, errors.New("cannot swap same currency")
	}

	// Check balance
	balance, err := s.BalanceService.GetBalance(walletID, fromCurrency)
	if err != nil {
		return nil, errors.New("wallet not found")
	}

	if balance.Amount < amount {
		return nil, errors.New("insufficient balance")
	}

	// Get FX rate
	rate, err := s.FXService.GetRate(fromCurrency, toCurrency)
	if err != nil {
		return nil, errors.New("failed to get FX rate")
	}

	// Calculate converted amount
	convertedAmount := amount * rate

	// Execute in transaction
	err = s.BalanceService.BalanceRepo.DB.Transaction(func(tx *gorm.DB) error {
		// Deduct from source - amount removed
		decreaseErr := s.BalanceService.UpdateInTx(tx, walletID, fromCurrency, -amount)
		if decreaseErr != nil {
			return decreaseErr
		}

		// Credit to target - convertedAmount added
		increaseErr := s.BalanceService.UpdateInTx(tx, walletID, toCurrency, convertedAmount)
		if increaseErr != nil {
			return increaseErr
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Create transaction record
	transaction := &models.Transaction{
		TxType:          "swap",
		SenderWalletID:  &walletID,
		FromCurrency:    fromCurrency,
		ToCurrency:      toCurrency,
		Amount:          amount,
		ConvertedAmount: convertedAmount,
		FxRate:          rate,
		Status: 		 "success",
		CreatedAt:       time.Now(),
	}

	// Save transaction
	if err := s.TransactionRepo.CreateTransaction(transaction); err != nil {
		return nil, err
	}

	return transaction, nil
}