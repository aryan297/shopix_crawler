# E-commerce Sitemap Crawler (Go) 

1. Project Overview
This repository contains the implementation of a web crawler that discovers product URLs across multiple e-commerce websites. The crawler uses Go and the Chromedp library to handle Single Page Applications (SPA) content and extract product URLs using dynamic page scraping. The crawler can scale to handle multiple websites and efficiently extract product links by analyzing HTML and matching URL patterns.

2. Tech Stack
Go (Golang): For the core crawler logic, providing a scalable and high-performance solution.

Chromedp: A Go library for controlling headless Chrome, useful for scraping dynamic content in SPAs.

utils: A custom utility package containing URL pattern matching to identify product pages.

3. Features
Crawl Multiple Domains: Capable of crawling a list of domains and discovering product URLs.

Dynamic Page Handling: Uses Chromedp to interact with JavaScript-heavy pages and extract product links.

Product URL Matching: Extracts product URLs based on common patterns (e.g., /product/, /p/, /item/).

Concurrency: The crawler can be extended to use concurrent or parallel crawling for scalability.

Headful Mode for Debugging: By default, the crawler runs in headless mode, but it can be switched to headful mode for debugging purposes to visualize the browser actions.

4. Setup Instructions
Prerequisites:
Go 1.18 or higher.

A working Go development environment.

Chrome browser or Chromium installed.

Installation:
Clone the repository:

bash
Copy
git clone https://github.com/yourusername/ecommerce-crawler.git
cd ecommerce-crawler
Install the required Go dependencies:

bash
Copy
go mod tidy
Download and install Chromedp: Chromedp automatically downloads the necessary browser binaries, but you can also configure it to use your local Chrome installation.

Build the project:

bash
Copy
go build -o crawler cmd/main.go
5. How it Works
The crawler works by visiting a series of e-commerce websites and identifying product URLs. It interacts with web pages using Chromedp to load and render JavaScript-heavy pages and extracts all links (<a> tags) from the page. These links are then filtered to identify potential product pages using a set of predefined URL patterns.

6. Crawling Process
Initiate Crawl: The crawler starts by visiting the base URL of the e-commerce site.

Page Load: It waits for the page to load and any JavaScript content to render.

Link Extraction: Using Chromedp, the crawler extracts all the URLs present in <a> tags on the page.

Pattern Matching: The links are filtered to match patterns indicative of product pages (e.g., /product/, /item/).

Result Collection: The valid product URLs are stored and can be saved in a file or database for further analysis.

7. Handling Dynamic Pages
Many modern e-commerce sites use Single Page Applications (SPAs), which load content dynamically via JavaScript. For this reason, traditional HTML scraping methods are insufficient. The crawler uses Chromedp to control a headless browser and simulate user interactions to render all JavaScript content. This allows the crawler to scrape content that wouldn't be accessible using static HTML parsing.

8. Known Issues and Troubleshooting
Slow Page Load: Some pages may take longer to load, especially with dynamic content. The crawler has a timeout set, but it may need to be adjusted based on site performance.

Incorrect URL Patterns: The crawler uses predefined patterns to match product URLs, but some e-commerce websites may use custom structures. You can extend the productPatterns list in the utils package to add more patterns as needed.

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
