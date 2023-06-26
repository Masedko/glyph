package middleware

import (
	"github.com/gofiber/fiber/v2"
	"tiktok-arena/internal/core/services"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	switch e := err.(type) {
	case services.ValidateError:
		code = fiber.StatusBadRequest
		message = e.Error()
	case services.RepositoryError:
		code = fiber.StatusBadRequest
		message = e.Error()
	case services.MatchAlreadyParsedError:
		code = fiber.StatusBadRequest
		message = e.Error()
	default:
		message = err.Error()
	}

	err = c.Status(code).JSON(fiber.Map{"message": message})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Cannot send error JSON message")
	}

	return nil
}
