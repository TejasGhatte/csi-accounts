package routes

import (
	"csi-accounts/internal/models"
	"csi-accounts/database"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func createResponseUser(userModel models.User) User {
	return User{
		ID:    userModel.ID,
		Name:  userModel.Name,
		Email: userModel.Email,
	}
}


func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}
	responseUser := createResponseUser(user)
	return c.Status(http.StatusOK).JSON(responseUser)
}
