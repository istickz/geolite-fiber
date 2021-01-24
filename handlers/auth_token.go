package handlers

import (
	"errors"
)

import (
	"github.com/gofiber/fiber/v2"
)

func CheckTokenIsValid(c *fiber.Ctx) error {
	a, err := TokenFromHeader(c, "Authorization", "Token")

	if err != nil {
		return err
	}

	if a == "123" {
		return nil
	}
	return errors.New("Missing or malformed TOKEN")
}

func TokenFromHeader(c *fiber.Ctx, header string, authScheme string) (string, error) {
	auth := c.Get(header)
	l := len(authScheme)
	if len(auth) > l+1 && auth[:l] == authScheme {
		return auth[l+1:], nil
	}
	return "", errors.New("Missing or malformed TOKEN")
}
