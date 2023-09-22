package api

import (
	"go-api/entity/message"
	"go-api/usecase/calculator"

	"github.com/gofiber/fiber/v2"
)

func GetMultiple(c *fiber.Ctx) error {

	messageData := new(message.ReceiveMessage)
	if err := c.BodyParser(messageData); err != nil {
		return c.SendString("Not valid")
	}

	result := calculator.GetValueMultiple(*messageData)
	return c.JSON(fiber.Map{"result": result})
}
