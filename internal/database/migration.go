package database

import (
	"log"

	"github.com/Zekeriyyah/stellar-x/internal/models"
)

// Run: applies database schema changes
func Run() {

	err := DB.AutoMigrate(
		&models.User{},
		&models.Balance{},
		&models.Wallet{},
		&models.AuditLog{},
		&models.Transaction{},

	)
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	log.Println("✅ Migrations completed successfully!")
}


