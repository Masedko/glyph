package routers

import (
	"github.com/gofiber/fiber/v2"
	"tiktok-arena/internal/api/controllers"
)

func NewGlyphRouter(c *controllers.GlyphController) func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Post("/:matchID", c.GetGlyphs)
	}
}
