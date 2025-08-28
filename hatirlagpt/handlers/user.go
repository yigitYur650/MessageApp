package handlers

import (
	"hatirlagpt/models"
	"hatirlagpt/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	var input RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Veri çözümlenemedi"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Şifreleme başarısız"})
	}

	newUser := models.User{
		ID:       uuid.New(),
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	storage.Users = append(storage.Users, newUser)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user_id":  newUser.ID,
		"username": newUser.Username,
	})
}

func Login(c *fiber.Ctx) error {
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Veri çözümlenemedi"})
	}

	for _, user := range storage.Users {
		if user.Email == input.Email {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
			if err == nil {
				return c.JSON(fiber.Map{
					"user_id":  user.ID,
					"username": user.Username,
				})
			}
			break
		}
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Email veya şifre yanlış"})
}
