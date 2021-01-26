package handlers

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
)

var db *geoip2.Reader

func init() {
	var err error
	db, err = geoip2.Open("geolite2_bases/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %v database.", db.Metadata().DatabaseType)
}

// GeoIP is a handler for IP address lookups
func GeoIP(c *fiber.Ctx, findType string) error {
	ipAddr := c.Params("ip", c.IP())

	// Check IP address format
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return c.Status(400).JSON(map[string]string{"status": "error", "message": "Invalid IP address"})
	}

	var record interface{}
	var err error

	// Switch type of return format
	switch findType {
	case "country":
		record, err = db.Country(ip)
	case "city":
		record, err = db.City(ip)
	default:
		record, err = db.Country(ip)
	}

	if err != nil {
		return err
	}

	return c.JSON(record)
}

func CloseDB() error {
	return db.Close()
}
