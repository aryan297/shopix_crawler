package crawler

import (
	"context"
	"fmt"
	"time"

	"github.com/aryan297/ecommerce-crawler-sitemap-go/internal/utils"
	"github.com/chromedp/chromedp"
)

// CrawlWithChromedp crawls a given URL, waits for the page to load, and extracts product links
func CrawlWithChromedp(startURL string, productPatterns []string) ([]string, error) {
	// Enable headful mode for debugging
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // See what's happening
		chromedp.Flag("disable-gpu", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Timeout for SPA rendering
	ctx, timeoutCancel := context.WithTimeout(ctx, 60*time.Second)
	defer timeoutCancel()

	var links []string

	// Navigate to the startURL and wait until the page content loads
	err := chromedp.Run(ctx,
		chromedp.Navigate(startURL),
		chromedp.WaitVisible("body", chromedp.ByQuery), // Wait for body to be visible
		chromedp.Sleep(5*time.Second),                  // Optional: allow extra time for SPAs to load
		chromedp.Evaluate(`Array.from(document.querySelectorAll("a"))
			.map(a => a.href)
			.filter(href => href && href.startsWith("http"))`, &links), // Filter only valid HTTP links
	)

	if err != nil {
		return nil, fmt.Errorf("navigation error: %w", err)
	}

	// Log the number of raw links found for debugging purposes
	fmt.Printf("Found %d raw links on %s\n", len(links), startURL)

	// Filter the links to match product patterns and remove duplicates
	return filterProductLinks(links, productPatterns), nil
}

// filterProductLinks filters out the product links that match specific patterns
func filterProductLinks(links []string, patterns []string) []string {
	uniqueLinks := make(map[string]struct{})
	var result []string

	for _, link := range links {
		if utils.MatchesAnyPattern(link, patterns) {
			if _, exists := uniqueLinks[link]; !exists {
				result = append(result, link)
				uniqueLinks[link] = struct{}{}
			}
		}
	}
	return result
}
