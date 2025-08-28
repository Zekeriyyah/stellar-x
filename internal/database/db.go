package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DATABASE_URL = "postgres://postgres:postgres@db:5432/stellar_x?sslmode=disable"
func InitDB() *gorm.DB {
	var dsn string
	// load .env file from root directory if not Render deployment environment
	if os.Getenv("RENDER") == "" {
		err := godotenv.Load()
		if err != nil {
			dsn = DATABASE_URL
		} else {
			dsn = os.Getenv("DATABASE_URL")
		}
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
