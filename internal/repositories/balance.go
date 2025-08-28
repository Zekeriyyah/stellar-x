package repositories

import (
	"github.com/Zekeriyyah/stellar-x/internal/models"
	"gorm.io/gorm"
)

type BalanceRepository struct {
	DB *gorm.DB
}

func NewBalanceRepository(db *gorm.DB) *BalanceRepository {
	return &BalanceRepository{DB: db}
}

// InitBalances creates zero balances for all supported stablecoins
func (r *BalanceRepository) InitBalances(balances []models.Balance) error {
	return r.DB.Create(&balances).Error
}

// FindByWalletIDAndCurrency retrieves a balance for a specific wallet and currency
func (r *BalanceRepository) FindByWalletIDAndCurrency(walletID uint, currency string) (*models.Balance, error) {
	var balance models.Balance
	err := r.DB.Where("wallet_id = ? AND currency = ?", walletID, currency).First(&balance).Error
	return &balance, err
}

// Update: updates a balance record (after deposit, swap, transfer)
func (r *BalanceRepository) Update(balance *models.Balance) error {
	return r.DB.Save(balance).Error
}

// UpdateInTx updates a balance within a transaction
func (r *BalanceRepository) UpdateInTx(tx *gorm.DB, balance *models.Balance) error {
	return tx.Save(balance).Error
}