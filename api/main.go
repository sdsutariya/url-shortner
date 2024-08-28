package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/sdsutariya/url-shortner/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"

		log.Fatal(app.Listen(":" + port))
	}
}
