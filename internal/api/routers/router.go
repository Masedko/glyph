package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "go-glyph-v2/docs"
)

func SetupRoutes(app *fiber.App,
	glyphRouter func(router fiber.Router)) {

	api := app.Group("/api")

	//	Use 'swag init' to generate new /docs files, details: https://github.com/gofiber/swagger#usage
	api.Get("/docs/*", swagger.HandlerDefault)
	// Redirect to docs
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/api/docs/")
	})

	api.Route("/glyph", glyphRouter)
}
