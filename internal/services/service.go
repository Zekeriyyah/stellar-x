package services

import (
	"github.com/Zekeriyyah/stellar-x/internal/models"
	"gorm.io/gorm"
)

type Services interface {
	//User services
	CreateUser(u *models.User) (int, string)
	Login(string, string) (string, int, string)
	GetUserByID(id uint) (*models.User, int, string)
	GetUserByEmail(email string) (*models.User, int, string) 

	// Wallet services
	CreateWalletWithBalances(email, label string) (*models.Wallet, error)
	GetWalletByUserID(userID uint) ([]models.Wallet, error)

	
	//audit services
	LogRequest(userID uint, walletID *uint, ip, device, browser, country, endpoint, method string) error
	GetAllLogs(userId uint) ([]models.AuditLog, error) 

	// balance services
	InitializeBalances(walletID uint) error
	GetBalance(walletID uint, currency string) (*models.Balance, error)
	Increase(walletID uint, currency string, amount float64) error
	Decrease(walletID uint, currency string, amount float64) error
	UpdateInTx(tx *gorm.DB, walletID uint, currency string, delta float64) error 

	//Deposit services
	Deposit(walletID uint, currency string, amount float64) (*models.Transaction, error)

	//FX Services
	GetRate(from, to string) (float64, error) 
	getFrankfurterRate(from, to string) (float64, error) 
	getCoinGeckoRate(from, to string) (float64, error)

	//AI Services
	Ask(query string) (string, error)

	//Swap Services
	Swap(walletID uint, fromCurrency, toCurrency string,	amount float64) (*models.Transaction, error) 

	//Transaction services
	GetByUserID(userID uint) ([]models.Transaction, error)
	GetRecent(page, pageSize int) (*PaginatedTransactions, error)

	//Transfer services
	Transfer(senderWalletID uint, receiverWalletID uint, fromCurrency string, toCurrency string, amount float64,) (*models.Transaction, error)

}