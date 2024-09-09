package main

import (
	"csi-accounts/config"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()
	app := fiber.New()
	log.Fatal(app.Listen(cfg.ServerPort))
}
