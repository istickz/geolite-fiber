package main

import (
	"geoip-maxmind/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())

	//app.Get("/geo/:ip?", handlers.GeoIP)

	app.Get("/geocountry/:ip?", func(c *fiber.Ctx) error {
		return handlers.GeoIP(c, "country")
	})

	app.Get("/geocity/:ip?", func(c *fiber.Ctx) error {
		return handlers.GeoIP(c, "city")
	})

	// Listen on port :3000
	log.Fatal(app.Listen(":3000"))
}
