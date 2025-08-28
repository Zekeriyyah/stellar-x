// internal/services/wallet_service.go
package services

import (
	"fmt"

	"github.com/Zekeriyyah/stellar-x/internal/models"
	"github.com/Zekeriyyah/stellar-x/internal/repositories"
)

type WalletService struct {
	WalletRepo *repositories.WalletRepository
	BalanceRepo *repositories.BalanceRepository
	UserRepo *repositories.UserRepository
}

func NewWalletService(w *repositories.WalletRepository, b *repositories.BalanceRepository, u *repositories.UserRepository) *WalletService {
	return &WalletService{
		WalletRepo:  w,
		BalanceRepo: b,
		UserRepo: u,
	}
}

// CreateWalletWithBalances creates wallet and initializes balances
func (s *WalletService) CreateWalletWithBalances(email, label string) (*models.Wallet, error) {
	
	// Get user with the email
	user, err:= s.UserRepo.FindUserByEmail(email)
	if err != nil {
		return &models.Wallet{}, fmt.Errorf("user not found %v", err)
	}
	
	wallet := &models.Wallet{
			UserID: user.ID,
			Label:  label,
		}
	wallet, err = s.WalletRepo.CreateWalletWithBalanceTx(wallet.UserID, wallet.Label)
	if err != nil {
		return &models.Wallet{}, err
	}
	return wallet, nil
}

// GetWalletByUserID retrieves wallet via repository
func (s *WalletService) GetWalletByUserID(userID uint) (*models.Wallet, error) {
	return s.WalletRepo.GetWalletByUserID(userID)
}
