package response

import "github.com/gofiber/fiber/v2"

func MessageResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
