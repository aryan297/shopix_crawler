# E-commerce Sitemap Crawler (Go)

## ✅ Features

- Crawls full webpages using headless Chrome (`chromedp`)
- Handles JavaScript-heavy Single Page Applications (SPAs)
- Extracts product/category links using domain-specific and generic patterns
- Outputs clean, deduplicated, filter-matched URLs
- Easily extendable with custom rules per domain

## ▶️ Run

```bash
go run cmd/crawler/main.go
```

## 🧪 Output

Results saved to `output/results.json` in the format:

```json
{
  "https://www.virgio.com/": ["https://www.virgio.com/product/123"],
  ...
}
```

## 🔗 Domains Tested

- https://www.virgio.com/
- https://www.tatacliq.com/
- https://nykaafashion.com/
- https://www.westside.com/
