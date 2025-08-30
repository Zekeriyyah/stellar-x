package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func InitDB() *gorm.DB {
	var dsn string

	// either using postgres on render or from docker
	if os.Getenv("RENDER") == "true" {
		dsn = os.Getenv("DATABASE_URL")
	} else {
		dsn = "postgres://postgres:postgres@db:5432/stellar_x?sslmode=disable"
	}	

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}

	DB = db
	log.Println("✅ Database connected successfully!")

	//Auto-migrate the db
	Run()
	return DB
}
