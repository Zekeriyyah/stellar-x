package repositories

import (
	"github.com/Zekeriyyah/stellar-x/internal/models"
	"gorm.io/gorm"
)


type TransactionRepository struct {
	DB *gorm.DB
}

type IDonly struct {
	ID uint
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) CreateTransaction(txn *models.Transaction) error {
	return r.DB.Create(txn).Error
}

// GetTransactionsByWalletID retrieves all transactions for a wallet
func (r *TransactionRepository) GetTransactionsByWalletID(walletID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.DB.Where(
		"sender_wallet_id = ? OR receiver_wallet_id = ?", walletID, walletID,
	).Order("created_at DESC").Find(&transactions).Error
	return transactions, err
}

// GetTransactionsByUserID uses wallet lookup
func (r *TransactionRepository) GetTransactionsByUserID(userID uint) ([]models.Transaction, error) {
	WalletIDs := []IDonly{}
	err := r.DB.Model(&models.Wallet{}).Select("id").Where("user_id = ?", userID).Find(&WalletIDs).Error
	if err != nil {
		return nil, err
	}
	
	transactions := []models.Transaction{}
	for _, walletId := range WalletIDs {

		tx, err := r.GetTransactionsByWalletID(walletId.ID)
		if err != nil {
			return transactions, err
		}
		transactions = append(transactions, tx...)
	}

	return transactions, nil
}

// GetRecent retrieves paginated recent transactions
func (r *TransactionRepository) GetRecent(limit, offset int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.DB.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&transactions).Error
	return transactions, err
}

// GetCount returns total number of transactions
func (r *TransactionRepository) GetCount() (int64, error) {
	var count int64
	err := r.DB.Model(&models.Transaction{}).Count(&count).Error
	return count, err
}