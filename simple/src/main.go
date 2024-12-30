package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())
	app.Use(cache.New(cache.Config{
		Expiration:   24 * time.Hour,
		CacheControl: true,
	}))

	app.Static("/favicon.ico", "./static/images/favicon.png", fiber.Static{MaxAge: 3600})
	app.Static("/static", "./static", fiber.Static{MaxAge: 3600})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{}, "layouts/main")
	})

	app.Get("/about.html", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{}, "layouts/main")
	})

	log.Fatal(app.Listen(":80"))
}
