package repositories

import (
	"context"

	"github.com/Zekeriyyah/stellar-x/internal/models"
)



type Repository interface {
	CreateUser(*models.User) error
	FindUserByEmail(string) (*models.User, error)
	FindUserByID(uint) (*models.User, error)

	// Wallet repo
	CreateWallet(wallet *models.Wallet) error
	InitBalances(balances []models.Balance) error
	GetWalletByUserID(userID uint) (*models.Wallet, error)
	CreateWalletWithBalanceTx(context.Context, string, string) (error, *models.User)
}
