package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-glyph-v2/configuration"
	"go-glyph-v2/internal/api/controllers"
	"go-glyph-v2/internal/api/middleware"
	"go-glyph-v2/internal/api/routers"
	"go-glyph-v2/internal/core/services"
	"go-glyph-v2/internal/data/database"
	"go-glyph-v2/internal/data/repository"
	"log"
)

// @title			Glyph Dota 2 REST API
// @version		1.0
// @description	API for Glyph Dota 2 application
// @host			go-glyph-v2-f53b68856ba5.herokuapp.com
// @BasePath		/api
func Run(c *configuration.EnvConfigModel) {

	db := database.ConnectDB(c)

	glyphRepository := repository.NewGlyphRepository(db)

	glyphService := services.NewGlyphService(glyphRepository)
	// stratzService := services.NewStratzService(c.STRATZToken)
	// opendotaService := services.NewOpendotaService()
	goSteamService := services.NewGoSteamService(c.SteamLoginUsername, c.SteamLoginPassword, c.SteamTwoFactorCode, c.SteamAuthCode)
	valveService := services.NewValveService()
	mantaService := services.NewMantaService()

	glyphController := controllers.NewGlyphController(glyphService, goSteamService,
		// opendotaService, stratzService,
		valveService, mantaService)

	glyphRouter := routers.NewGlyphRouter(glyphController)

	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})

	//	Logger middleware for logging HTTP request/response details
	app.Use(logger.New())

	//	CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://s3rbug.github.io",
		AllowHeaders: "POST",
	}))

	routers.SetupRoutes(app, glyphRouter)

	port := c.Port
	if port == "" {
		port = "8000"
	}

	log.Fatal(app.Listen(":" + port))
}
