# Content Management Improvements

Patterns borrowed from Astro and Hugo to make this repo easier to work with.

## Current Pain Points

- Posts are manually linked in `notes.html` and routes are hardcoded in `main.go`
- `homelab.md` has a `Last updated:` line that looks like frontmatter but isn't parsed — Goldmark renders it as plain text
- No content discovery, no structured metadata

## Patterns to Adopt

### 1. Real Frontmatter Parsing

Use YAML frontmatter (`---` delimited) to attach metadata to content files. A Go library like `adrg/frontmatter` can extract it before passing the body to Goldmark.

```yaml
---
title: Homelab
date: 2025-01-04
slug: homelab
draft: false
tags: [homelab, infrastructure]
---
```

This is the foundation for everything else below.

### 2. Auto-Discovery of Content

Instead of hardcoding links in `notes.html`, scan `static/content/` at startup, read each `.md` file's frontmatter, and build a sorted list automatically. New post = new `.md` file, nothing else to touch.

In the handler, this would be a function that `os.ReadDir("static/content/")`, parses frontmatter from each file, sorts by date, filters drafts, and passes the list to the `notes.html` template.

### 3. Draft Support

A `draft: true` field in frontmatter. Posts with that flag are excluded from production but visible in dev. Filter based on an env var or build flag.

### 4. Slug-Based Routing

Instead of `/content/:fileName` exposing the raw filename, derive the URL from a `slug` field in frontmatter (or auto-generate from the title). Gives cleaner URLs like `/notes/homelab` and decouples the URL from the filename.

### 5. Date-Based Sorting + Display

Once `date` is parsed from frontmatter, the notes index page can auto-sort newest-first and display dates next to titles. No manual ordering.

### 6. Layout Selection via Frontmatter

Let content specify a layout (e.g. `layout: wide`) in frontmatter to pick a different template. Currently there's only `layouts/markdown.html` — this becomes useful as more layouts are added.

## Suggested Priority

1. **Frontmatter parsing** — foundation for everything else
2. **Auto-discovery in `notes.html`** — eliminates manually adding new posts
3. **Draft support** — commit WIP without publishing

These three together give the core Hugo/Astro workflow: create a `.md` file and it shows up on the site. Tags, RSS, and slug routing layer on naturally after.
