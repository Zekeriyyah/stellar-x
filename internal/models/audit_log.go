package models

import "time"

type AuditLog struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    UserID    uint      `gorm:"not null" json:"user_id"`
    WalletID  *uint     `json:"wallet_id,omitempty"` // Optional
    IPAddress string    `json:"ip_address"`
    Device    string    `json:"device"`
    Browser   string    `json:"browser"`
    Country   string    `json:"country"`
    Path      string    `json:"path"`
    Method    string    `json:"method"`
    CreatedAt time.Time `json:"created_at"`

    // Relations
    User   User    `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL" json:"user"`
    Wallet *Wallet `gorm:"foreignKey:WalletID;constraint:OnDelete:SET NULL" json:"wallet,omitempty"`
}
