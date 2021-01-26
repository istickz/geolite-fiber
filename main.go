package main

import (
	"fmt"
	"geoip-maxmind/handlers"
	"geoip-maxmind/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	//log.Fatal(app.Listen(":3000"))

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)   // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")
	handlers.CloseDB()
}
