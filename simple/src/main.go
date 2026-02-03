package main

import (
	"crypto/tls"
	"html/template"
	"log"
	"path/filepath"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/valyala/fasthttp"

	"github.com/enmyj/ianmyjerdotcom/handlers"
)

var honeypotPaths = regexp.MustCompile(`^/(\.env|\.git|\.aws|\.ssh|wp-login\.php|wp-admin|xmlrpc\.php|phpmyadmin|pma|admin|administrator|config\.php|backup\.sql|dump\.sql|actuator|debug|trace|server-status|jenkins|cgi-bin|shell|eval-stdin\.php|training.*)`)
var honeypotClient = &fasthttp.Client{
	TLSConfig:                     &tls.Config{InsecureSkipVerify: true},
	NoDefaultUserAgentHeader:      true,
	DisableHeaderNamesNormalizing: true,
	DisablePathNormalizing:        true,
	ReadTimeout:                   30 * time.Second,
	WriteTimeout:                  30 * time.Second,
}

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
		Next: func(c *fiber.Ctx) bool {
			return honeypotPaths.MatchString(c.Path())
		},
	}))

	app.Static("/favicon.ico", "./static/images/favicon.png", fiber.Static{MaxAge: 604800})
	app.Static("/robots.txt", "./static/robots.txt", fiber.Static{MaxAge: 604800})
	app.Static("/static", "./static", fiber.Static{MaxAge: 604800})

	app.Use(func(c *fiber.Ctx) error {
		if honeypotPaths.MatchString(c.Path()) {

			req := fasthttp.AcquireRequest()
			resp := fasthttp.AcquireResponse()
			defer fasthttp.ReleaseRequest(req)
			defer fasthttp.ReleaseResponse(resp)

			req.SetRequestURI("https://rnsaffn.com/poison2/")
			req.Header.SetMethod(c.Method())
			c.Request().Header.VisitAll(func(key, value []byte) {
				req.Header.SetBytesKV(key, value)
			})
			req.SetBody(c.Body())

			if err := honeypotClient.Do(req, resp); err != nil {
				return err
			}

			// Copy all response headers
			resp.Header.VisitAll(func(key, value []byte) {
				c.Response().Header.SetBytesKV(key, value)
			})
			c.Status(resp.StatusCode())
			return c.Send(resp.Body())
		}
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{}, "layouts/main")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{}, "layouts/main")
	})

	app.Get("/notes", func(c *fiber.Ctx) error {
		return c.Render("notes", fiber.Map{}, "layouts/main")
	})

	contentDir, err := filepath.Abs("./static/content")
	if err != nil {
		log.Fatal(err)
	}
	app.Get("/content/:fileName", func(c *fiber.Ctx) error {
		return handlers.RenderMarkdown(c, contentDir)
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).Render("404", fiber.Map{}, "layouts/main")
	})

	log.Fatal(app.Listen(":8000"))
}
