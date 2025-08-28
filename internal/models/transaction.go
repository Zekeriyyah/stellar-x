package models

import "time"

type Transaction struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	TxType           string    `gorm:"not null" json:"tx_type"` // deposit, swap, transfer
	SenderWalletID   *uint     `json:"sender_wallet_id"`
	ReceiverWalletID *uint     `json:"receiver_wallet_id"`
	FromCurrency     string    `json:"from_currency"`
	ToCurrency       string    `json:"to_currency"`
	Amount           float64   `json:"amount"`
	ConvertedAmount  float64   `json:"converted_amount"`
	FxRate           float64   `json:"fx_rate"`
	Status 			 string	   `gorm:"not null;default:'completed'" json:"status"`
	TxHash 			 string    `gorm:"size:66;uniqueIndex" json:"tx_hash,omitempty"`
	CreatedAt        time.Time `json:"created_at"`

	// Relations
	SenderWallet   *Wallet `gorm:"foreignKey:SenderWalletID;constraint:OnDelete:SET NULL;" json:"-"`
	ReceiverWallet *Wallet `gorm:"foreignKey:ReceiverWalletID;constraint:OnDelete:SET NULL;" json:"-"`
}
