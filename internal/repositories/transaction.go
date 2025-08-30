package repositories

import (
	"github.com/Zekeriyyah/stellar-x/internal/models"
	"gorm.io/gorm"
)


type TransactionRepository struct {
	DB *gorm.DB
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
	var wallet models.Wallet
	err := r.DB.Select("id").First(&wallet, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return r.GetTransactionsByWalletID(wallet.ID)
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