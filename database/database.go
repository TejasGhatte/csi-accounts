package database

import (
	"csi-accounts/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=IIMB@1802 dbname=csi_rbac_system port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	// Open the connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	DB = db

	DB.AutoMigrate(&models.User{})
	log.Println("Database connected successfully!")
}
