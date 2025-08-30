package services

import (
	"errors"
	"time"

	"github.com/Zekeriyyah/stellar-x/internal/models"
	"github.com/Zekeriyyah/stellar-x/internal/repositories"
	"gorm.io/gorm"
)

type TransferService struct {
	WalletRepo      *repositories.WalletRepository
	BalanceService  *BalanceService
	TransactionRepo *repositories.TransactionRepository
	FXService       *FXService
}

func NewTransferService(w *repositories.WalletRepository,	b *BalanceService,	tx *repositories.TransactionRepository, fx *FXService,) *TransferService {
	return &TransferService{
		WalletRepo:      w,
		BalanceService:  b,
		TransactionRepo: tx,
		FXService:       fx,
	}
}

// Transfer sends funds to another wallet
// If currencies differ, auto-convert using FX rate
func (s *TransferService) Transfer(senderWalletID uint, receiverWalletID uint, fromCurrency string, toCurrency string, amount float64,) (*models.Transaction, error) {
	// Validate input - no negative input
	if amount <= 0 {
		return nil, errors.New("amount must be positive")
	}

	// Check sender balance - wallet exist and enough money
	senderBalance, err := s.BalanceService.GetBalance(senderWalletID, fromCurrency)
	if err != nil {
		return nil, errors.New("sender wallet not found")
	}
	if senderBalance.Amount < amount {
		return nil, errors.New("insufficient balance")
	}

	// Check if reciever wallet exist
	_, err = s.WalletRepo.GetWalletByUserID(receiverWalletID)
	if err != nil {
		return nil, errors.New("receiver wallet not found")
	}

	// Execute in transaction
	var convertedAmount float64
	var fxRate float64

	err = s.BalanceService.BalanceRepo.DB.Transaction(func(tx *gorm.DB) error {
		// If same currency, direct transfer
		if fromCurrency == toCurrency {
			// Deduct from sender
			decreaseErr := s.BalanceService.UpdateInTx(tx, senderWalletID, fromCurrency, -amount)
			if decreaseErr != nil {
				return decreaseErr
			}

			// Credit to receiver
			increaseErr := s.BalanceService.UpdateInTx(tx, receiverWalletID, toCurrency, amount)
			if increaseErr != nil {
				return increaseErr
			}

			convertedAmount = amount
			fxRate = 1.0
		} else {
			// Different currency → use FX rate
			rate, err := s.FXService.GetRate(fromCurrency, toCurrency)
			if err != nil {
				return errors.New("failed to get FX rate")
			}

			convertedAmount = amount * rate
			fxRate = rate

			// Deduct from sender (in source currency)
			decreaseErr := s.BalanceService.UpdateInTx(tx, senderWalletID, fromCurrency, -amount)
			if decreaseErr != nil {
				return decreaseErr
			}

			// Credit to receiver (in target currency)
			increaseErr := s.BalanceService.UpdateInTx(tx, receiverWalletID, toCurrency, convertedAmount)
			if increaseErr != nil {
				return increaseErr
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Create transaction record
	transaction := &models.Transaction{
		TxType:           "transfer",
		SenderWalletID:   &senderWalletID,
		ReceiverWalletID: &receiverWalletID,
		FromCurrency:     fromCurrency,
		ToCurrency:       toCurrency,
		Amount:           amount,
		ConvertedAmount:  convertedAmount,
		FxRate:           fxRate,
		Status:			  "success" ,	
		CreatedAt:        time.Now(),
	}

	// Save transaction
	if err := s.TransactionRepo.CreateTransaction(transaction); err != nil {
		return nil, err
	}

	return transaction, nil
}