package tokenauth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)


var appToken string

func init()  {
	var exists bool

	appToken, exists = os.LookupEnv("APP_TOKEN")

	if exists != true {
		// Print the value of the environment variable
		log.Fatal("APP_TOKEN was not set")
	}
}

// New creates a new middleware handler
func New() fiber.Handler {

	// Return new handler
	return func(c *fiber.Ctx) error {
		err := CheckTokenIsValid(c)

		if err != nil {
			data := AuthError{err.Error()}
			return c.Status(fiber.StatusUnauthorized).JSON(data)
		}

		return c.Next()
	}
}

type AuthError struct {
	Error string
}

func TokenFromHeader(c *fiber.Ctx, header string, authScheme string) (string, error) {
	auth := c.Get(header)
	l := len(authScheme)
	if len(auth) > l+1 && auth[:l] == authScheme {
		return auth[l+1:], nil
	}
	return "", errors.New("Missing or malformed TOKEN")
}

func CheckTokenIsValid(c *fiber.Ctx) error {
	a, err := TokenFromHeader(c, "Authorization", "Token")

	if err != nil {
		return err
	}

	if a == appToken {
		return nil
	}
	return errors.New("Bad TOKEN")
}