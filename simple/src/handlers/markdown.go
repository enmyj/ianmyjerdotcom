package handlers

import (
	"bytes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/yuin/goldmark"
)

func RenderMarkdown(c *fiber.Ctx) error {
	contentPath := c.Params("*")
	content, err := os.ReadFile("./static/content/" + contentPath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error reading Markdown file")
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(content, &buf); err != nil {
		panic(err)
	}

	return c.Render("markdown", fiber.Map{
		"Content": buf.String(),
	}, "layouts/main")
}
