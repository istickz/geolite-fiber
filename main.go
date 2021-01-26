package main

import (
	"geoip-maxmind/handlers"
	"geoip-maxmind/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(tokenauth.New())

	app.Get("/geocountry/:ip?", func(c *fiber.Ctx) error {
		return handlers.GeoIP(c, "country")
	})

	app.Get("/geocity/:ip?", func(c *fiber.Ctx) error {
		return handlers.GeoIP(c, "city")
	})

	// Listen on port :3000
	log.Fatal(app.Listen(":3000"))
}
