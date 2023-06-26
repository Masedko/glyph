package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
	"tiktok-arena/configuration"
	"tiktok-arena/internal/api/controllers"
	"tiktok-arena/internal/api/middleware"
	"tiktok-arena/internal/api/routers"
	"tiktok-arena/internal/core/services"
	"tiktok-arena/internal/data/database"
	"tiktok-arena/internal/data/repository"
)

//	@title			Glyph Dota 2 REST API
//	@version		1.0
//	@description	API for Glyph Dota 2 application
//	@host			localhost:8000
//	@BasePath		/api
func Run(c *configuration.EnvConfigModel) {

	db := database.ConnectDB(c)

	glyphRepository := repository.NewGlyphRepository(db)

	glyphService := services.NewGlyphService(glyphRepository)
	stratzService := services.NewStratzService(c.STRATZToken)
	valveService := services.NewValveService()
	mantaService := services.NewMantaService()

	glyphController := controllers.NewGlyphController(glyphService, stratzService, valveService, mantaService)

	glyphRouter := routers.NewGlyphRouter(glyphController)

	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})

	//	Logger middleware for logging HTTP request/response details
	app.Use(logger.New())

	//	CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	routers.SetupRoutes(app, glyphRouter)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Fatal(app.Listen(":" + port))
}
