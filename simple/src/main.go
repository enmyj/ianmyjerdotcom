package main

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/yuin/goldmark"
)

func main() {
	engine := html.New("./views", ".html")
	engine.AddFunc(
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())
	app.Use(cache.New(cache.Config{
		Expiration:   24 * time.Hour,
		CacheControl: true,
	}))

	app.Static("/favicon.ico", "./static/images/favicon.png", fiber.Static{MaxAge: 604800})
	app.Static("/robots.txt", "./static/robots.txt", fiber.Static{MaxAge: 604800})
	app.Static("/static", "./static", fiber.Static{MaxAge: 604800})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{}, "layouts/main")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{}, "layouts/main")
	})

	app.Get("/markdown", func(c *fiber.Ctx) error {
		content, err := os.ReadFile("./static/content/test.md")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error reading Markdown file")
		}

		var buf bytes.Buffer
		if err := goldmark.Convert(content, &buf); err != nil {
			panic(err)
		}

		return c.Render("ian", fiber.Map{
			"Content": buf.String(),
		}, "layouts/main")
	})

	log.Fatal(app.Listen(":8000"))
}
