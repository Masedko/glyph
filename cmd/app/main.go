package main

import (
	"log"
	"tiktok-arena/configuration"
	"tiktok-arena/internal/api/app"
)

func main() {
	err := configuration.LoadConfig(".env")
	if err != nil {
		log.Fatalln("Failed to load environment variables!", err.Error())
	}
	app.Run(&configuration.EnvConfig)
}
