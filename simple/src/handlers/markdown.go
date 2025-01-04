package handlers

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/yuin/goldmark"
)

func RenderMarkdown(c *fiber.Ctx, contentDir string) error {
	fileName := c.Params("fileName")
	absFullPath := filepath.Join(contentDir, filepath.Clean(fileName))

	if !strings.HasPrefix(absFullPath, contentDir) {
		return c.Status(fiber.StatusForbidden).SendString("Invalid file path")
	}

	content, err := os.ReadFile(absFullPath)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("File not found")
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(content, &buf); err != nil {
		panic(err)
	}

	return c.Render("layouts/markdown", fiber.Map{
		"Content": buf.String(),
	}, "layouts/main")
}
