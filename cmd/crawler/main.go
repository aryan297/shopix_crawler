package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aryan297/ecommerce-crawler-sitemap-go/internal/crawler"
	"github.com/aryan297/ecommerce-crawler-sitemap-go/internal/utils"
)

func main() {
	domains := []string{
		"https://www.virgio.com/",
		"https://www.tatacliq.com/",
		"https://nykaafashion.com/",
		"https://www.westside.com/",
	}

	results := make(map[string][]string)

	for _, domain := range domains {
		fmt.Println("🔍 Crawling:", domain)

		// Dynamically get domain-specific patterns
		patterns := utils.GetPatternsForURL(domain)

		urls, err := crawler.CrawlWithChromedp(domain, patterns)
		if err != nil {
			log.Printf("❌ Error crawling %s: %v\n", domain, err)
			continue
		}

		results[domain] = urls
		fmt.Printf("✅ %d matched links for %s\n\n", len(urls), domain)
	}

	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Fatalf("❌ Failed to marshal results: %v", err)
	}

	err = os.WriteFile("output/results.json", data, 0644)
	if err != nil {
		log.Fatalf("❌ Failed to write results file: %v", err)
	}

	fmt.Println("🚀 Crawling complete. Results written to output/results.json")
}
