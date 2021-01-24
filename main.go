package main

import (
	"geoip-maxmind/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//app.Get("/geo/:ip?", handlers.GeoIP)

	app.Get("/geocountry/:ip?", func(c *fiber.Ctx) error {
		handlers.GeoIP(c, "country")
		return nil
	})

	app.Get("/geocity/:ip?", func(c *fiber.Ctx) error {
		handlers.GeoIP(c, "city")
		return nil
	})

	// Listen on port :3000
	log.Fatal(app.Listen(":3000"))
}
