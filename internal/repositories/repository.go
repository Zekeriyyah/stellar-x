package repositories

import (
	"github.com/Zekeriyyah/stellar-x/internal/models"
	"gorm.io/gorm"
)



type Repository interface {
	// User repo
	CreateUser(*models.User) error
	FindUserByEmail(string) (*models.User, error)
	FindUserByID(uint) (*models.User, error)

	// Wallet repo
	CreateWallet(wallet *models.Wallet) error
	InitBalances(balances []models.Balance) error
	GetWalletByUserID(userID uint) (*models.Wallet, error)
	CreateWalletWithBalanceTx(userId uint, label string)  (*models.Wallet, error)

	// Balance repo
	//InitBalances(balances []models.Balance) error
	FindByWalletIDAndCurrency(walletID uint, currency string) (*models.Balance, error)
	Update(balance *models.Balance) error 
	UpdateInTx(tx *gorm.DB, balance *models.Balance) error

	// Audit repo
	Create(log *models.AuditLog) error
	GetAuditLogByUserID(userID uint) ([]models.AuditLog, error) 

	// Transaction repo
	GetTransactionsByWalletID(walletID uint) ([]models.Transaction, error)
	GetTransactionsByUserID(userID uint) ([]models.Transaction, error)
	GetRecent(limit, offset int) ([]models.Transaction, error)
	GetCount() (int64, error)
}
