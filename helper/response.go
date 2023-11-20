package helper

import "github.com/gofiber/fiber/v2"

func SendResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": statusCode > 200 && statusCode < 300,
		"message": message,
		"data":    data,
	})
}
