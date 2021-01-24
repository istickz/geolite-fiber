package handlers

import (
	"fmt"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/oschwald/maxminddb-golang"
)

// See https://pkg.go.dev/github.com/oschwald/geoip2-golang#City for a full list of options you can use here to modify
// what data is returned for a specific IP.
type ipCountryLookup struct {
	Country struct {
		IsoCode string `maxminddb:"iso_code"`
	} `maxminddb:"country"`
}

type ipCityLookup struct {
	City struct {
		GeoNameID uint              `maxminddb:"geoname_id"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"city"`
	Country struct {
		IsoCode string `maxminddb:"iso_code"`
	} `maxminddb:"country"`
	Location struct {
		AccuracyRadius uint16 `maxminddb:"accuracy_radius"`
	} `maxminddb:"location"`
}

var geoIPCountryDb *maxminddb.Reader
var geoIPCityDb *maxminddb.Reader

func init() {
	// Load MaxMind DB
	var err error
	geoIPCountryDb, err = maxminddb.Open("geolite2_bases/GeoLite2-Country.mmdb")
	if err != nil {
		fmt.Println("Unable to load 'GeoLite2-Country.mmdb'.")
		panic(err)
	}

	geoIPCityDb, err = maxminddb.Open("geolite2_bases/GeoLite2-City.mmdb")

	if err != nil {
		fmt.Println("Unable to load 'GeoLite2-City.mmdb'.")
		panic(err)
	}
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

	switch findType {
	case "country":
		record, err = findCountry(ip)
	case "city":
		record, err = findCity(ip)
	default:
		record, err = findCountry(ip)
	}

	if err != nil {
		return err
	}

	return c.JSON(record)
}

func findCountry(ip net.IP) (record interface{}, err error)  {
	record = new(ipCountryLookup)
	err = geoIPCountryDb.Lookup(ip, &record)
	if err != nil {
		return err, nil
	}
	return record, nil
}

func findCity(ip net.IP) (record interface{}, err error)  {
	record = new(ipCityLookup)
	err = geoIPCityDb.Lookup(ip, &record)
	if err != nil {
		return err, nil
	}
	return record, nil
}

