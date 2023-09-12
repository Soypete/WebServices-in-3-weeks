package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Route("/3weeks-go", func(api fiber.Router) {
		api.Get("/testGet", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		}).Name("foo") // /test/foo (name: test.foo)
		api.Post("/testPost", func(c *fiber.Ctx) error {
			body := c.Body()
			if len(body) == 0 {
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return c.SendString(string(body))
		}).Name("bar") // /test/bar (name: test.bar)
	}, "test.")

	app.Listen(":3000")
}
