package main

import (
	"go-glyph-v2/configuration"
	"go-glyph-v2/internal/api/app"
	"log"
)

func main() {
	err := configuration.LoadConfig(".env")
	if err != nil {
		log.Fatalln("Failed to load environment variables!", err.Error())
	}
	app.Run(&configuration.EnvConfig)
}
