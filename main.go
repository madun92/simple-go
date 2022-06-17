package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("html", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "Hello World",
		})
	})

	app.Listen(":3000")
}
