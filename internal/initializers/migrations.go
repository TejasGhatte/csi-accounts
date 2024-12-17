package initializers

import (
	"log"
	"fmt"

	"csi-accounts/internal/models"
)

func RunMigrations() {
	if DB == nil {
		fmt.Println("Database not connected")
	}

    err := DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.Event{},
		&models.EventMembership{},
		&models.Scope{},
		&models.Client{},
		&models.ClientScope{},
		&models.UserScope{},
	)

	if err != nil {
		log.Fatal("Failed to run migrations")
	}
}