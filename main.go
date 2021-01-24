package main

import (
	"fmt"
	"geoip-maxmind/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"log"
)



type ApiError struct {
	Message string
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())


	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("🥇 First handler")
        err := handlers.CheckTokenIsValid(c)

        if err != nil {
			fmt.Println("🥇 send error")

			c.Status(fiber.StatusUnauthorized)
			data := ApiError{err.Error()}

			return c.JSON(data)
		}

		return c.Next()
	})

	app.Get("/geocountry/:ip?", func(c *fiber.Ctx) error {
		return handlers.GeoIP(c, "country")
	})

	app.Get("/geocity/:ip?", func(c *fiber.Ctx) error {
		return handlers.GeoIP(c, "city")
	})

	// Listen on port :3000
	log.Fatal(app.Listen(":3000"))
}
