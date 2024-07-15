package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	urlShortener := NewUrlShortener()

	app.Post("/api/v1/shortener", func(c *fiber.Ctx) error {
		data := make(map[string]string)
		if err := c.BodyParser(&data); err != nil {
			return err
		}

		url, ok := data["url"]

		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "url is required"})
		}

		url = adjustHTTPS(url)
		fmt.Println(url)

		if !isURLValid(url) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "url is not valid"})
		}

		serverOrigin := c.Hostname()
		return c.JSON(map[string]string{"origin": serverOrigin, "short_url": urlShortener.GetShortenUrl(url)})
	})

	app.Get("/:route", func(c *fiber.Ctx) error {
		route := c.Params("route")
		redirectUrl, error := urlShortener.GetRedirectUrl(route)
		if error != nil {
			return c.JSON(map[string]string{"error": "redirect url is invalid"})
		}
		return c.Redirect(redirectUrl, fiber.StatusPermanentRedirect)
	})

	log.Fatal(app.Listen(":3000"))
}
