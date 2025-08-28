package services

import (
	"context"

	"github.com/Zekeriyyah/stellar-x/internal/models"
)

type Services interface {
	CreateUser(u *models.User) (int, string)
	Login(string, string) (string, int, string)

	// Wallet services
	CreateWalletWithBalances(context.Context, string, string) (error, *models.User)
	// GetWallet(userId string) (*models.Wallet, error)

	
	
}