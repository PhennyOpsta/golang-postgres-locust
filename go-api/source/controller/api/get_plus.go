package api

import (
	"go-api/entity/message"
	"go-api/usecase/calculator"

	"github.com/gofiber/fiber/v2"
)

func GetPlus(c *fiber.Ctx) error {

	messageData := new(message.ReceiveMessage)
	if err := c.BodyParser(messageData); err != nil {
		return c.SendString("Not valid")
	}

	result := calculator.GetValuePlus(*messageData)
	return c.JSON(fiber.Map{"result": result})
}
