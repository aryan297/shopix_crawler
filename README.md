# E-commerce Sitemap Crawler (Go)

## 1. Project Overview
This repository contains the implementation of a web crawler that discovers product URLs across multiple e-commerce websites. The crawler uses **Go** and the **Chromedp** library to handle Single Page Applications (SPA) content and extract product URLs using dynamic page scraping. The crawler can scale to handle multiple websites and efficiently extract product links by analyzing HTML and matching URL patterns.

## 2. Tech Stack
- **Go (Golang)**: For the core crawler logic, providing a scalable and high-performance solution.
- **Chromedp**: A Go library for controlling headless Chrome, useful for scraping dynamic content in SPAs.
- **utils**: A custom utility package containing URL pattern matching to identify product pages.

## 3. Features
- **Crawl Multiple Domains**: Capable of crawling a list of domains and discovering product URLs.
- **Dynamic Page Handling**: Uses **Chromedp** to interact with JavaScript-heavy pages and extract product links.
- **Product URL Matching**: Extracts product URLs based on common patterns (e.g., `/product/`, `/p/`, `/item/`).
- **Concurrency**: The crawler can be extended to use concurrent or parallel crawling for scalability.
- **Headful Mode for Debugging**: By default, the crawler runs in headless mode, but it can be switched to headful mode for debugging purposes to visualize the browser actions.

## 4. Setup Instructions
### Prerequisites:
- **Go 1.18** or higher.
- A working Go development environment.
- **Chrome** browser or **Chromium** installed.

- **1git clone git@github.com:aryan297/shopix_crawler.git

cd /shopix_crawler

--2. go mod tidy

--3. go run cmd/crawler/main.go 

