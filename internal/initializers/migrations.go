package initializers

import (
	"log"
	"fmt"
)

func RunMigrations() {
	if DB == nil {
		fmt.Println("Database not connected")
	}

    err := DB.AutoMigrate()

	if err != nil {
		log.Fatal("Failed to run migrations")
	}
}