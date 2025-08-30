package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func InitDB() *gorm.DB {
	var dsn string

	// either using postgres on render or from docker
	if os.Getenv("RENDER") == "true" {
		
		err := godotenv.Load()
		if err != nil {
			log.Fatal("failed to load .env file")
		}

		dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
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
