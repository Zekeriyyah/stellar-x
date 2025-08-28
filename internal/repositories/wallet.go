package repositories

import (
	"github.com/Zekeriyyah/stellar-x/internal/models"
	"gorm.io/gorm"
)

type WalletRepository struct {
	DB *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{
		DB: db,
	}
}

func (r *WalletRepository) CreateWallet(wallet *models.Wallet) error {
	return r.DB.Create(wallet).Error
}

func (r *WalletRepository) InitBalances(balances []models.Balance) error {
	return r.DB.Create(&balances).Error
}

func (r *WalletRepository) CreateWalletWithBalanceTx(userId uint, label string)  (*models.Wallet, error) {
	wallet := models.Wallet{
		UserID: userId,
		Label: label,
	}

	balances := []models.Balance{}

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		
		if err := tx.Create(&wallet).Error; err != nil {
			return err
		}

		balances = []models.Balance{
			{WalletID: wallet.ID, Currency: "cNGN", Amount: 0},
			{WalletID: wallet.ID, Currency: "cXAF", Amount: 0},
			{WalletID: wallet.ID, Currency: "USDx", Amount: 0},
			{WalletID: wallet.ID, Currency: "EURx", Amount: 0},
		}
		return tx.Create(&balances).Error
	})
	if err != nil {
		return &models.Wallet{}, err
	}
	wallet.Balances = balances
	return &wallet, nil
}

func (r *WalletRepository) GetWalletByUserID(userID uint) (*models.Wallet, error) {
	wallet := models.Wallet{}
	err := r.DB.Preload("Balances").First(&wallet, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}
