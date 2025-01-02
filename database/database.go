package database

import (
	"csi-accounts/internal/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func LoadEnvVariables() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}


func ConnectDatabase() {
	
	LoadEnvVariables()

	
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)

	// Connect to the database using the DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	DB = db
	DB.AutoMigrate(&models.User{}, &models.Event{}, &models.EventMembership{}) // Include all models you need

	log.Println("Database connected successfully!")
}
