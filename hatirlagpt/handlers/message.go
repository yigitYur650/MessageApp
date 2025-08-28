// handlers/message.go

package handlers

import (
	"time"

	"hatirlagpt/models"
	"hatirlagpt/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Mesaj gönderme endpoint'i
func SendMessage(c *fiber.Ctx) error {
	type MessageInput struct {
		SenderID         uuid.UUID `json:"sender_id"`
		ReceiverUsername string    `json:"receiver_username"`
		Content          string    `json:"content"`
	}

	var input MessageInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Veri çözümlenemedi"})
	}

	var receiverUUID uuid.UUID
	found := false
	for _, u := range storage.Users {
		if u.Username == input.ReceiverUsername {
			receiverUUID = u.ID
			found = true
			break
		}
	}

	if !found {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Alıcı bulunamadı"})
	}

	newMessage := models.Message{
		ID:         uuid.New(),
		SenderID:   input.SenderID,
		ReceiverID: receiverUUID,
		Content:    input.Content,
		Timestamp:  time.Now(),
	}

	storage.Messages = append(storage.Messages, newMessage)

	return c.Status(fiber.StatusCreated).JSON(newMessage)
}
