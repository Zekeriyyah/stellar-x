package models

import (
	"time"
)

type Wallet struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Label     string    `json:"label"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	User     User      `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	Balances []Balance `gorm:"foreignKey:WalletID" json:"balances,omitempty"`
}


type Balance struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	WalletID  uint      `gorm:"not null" json:"wallet_id"`
	Currency  string    `gorm:"size:4;not null" json:"currency"` // e.g. USD, NGN, EUR
	Amount    float64   `gorm:"not null;default:0" json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Wallet *Wallet `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
}

