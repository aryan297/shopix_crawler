package utils

import (
	"net/url"
	"regexp"
	"strings"
)

// Fallback product patterns for general sites
var GenericPatterns = []string{
	`/product/[\w-]+`,
	`/products/[\w-]+`,
	`/collections/[\w-]+`,
	`/collections/[\w-]+/[\w-]+`,
	`/pages/[\w-]+`,
	`/p/[\w-]+`,
	`/item/[\w-]+`,
	`/\d{5,}`, // numeric product IDs
}

// Domain-specific patterns
var DomainPatterns = map[string][]string{
	"tatacliq.com": {
		`/[^/]+/c-[\w\d]+`,          // e.g., /watches/c-msh15
		`/c-[\w\d]+`,                // e.g., /c-msh15
		`/[^/]+/c-[\w\d]+/page-\d+`, // paginated
	},
	"westside.com": {
		`/collections/[\w-]+`,
		`/products/[\w-]+`,
	},
	"nykaafashion.com": {
		`/product/[\w-]+`,
		`/p/[\w-]+`,
		`/item/[\w-]+`,
		`/\d{5,}`,
	},
	"virgio.com": {
		`/products/[\w-]+`,
	},
}

// Returns domain-specific patterns if available, else falls back to generic
func GetPatternsForURL(rawURL string) []string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return GenericPatterns
	}
	host := strings.TrimPrefix(u.Hostname(), "www.")
	if patterns, ok := DomainPatterns[host]; ok {
		return patterns
	}
	return GenericPatterns
}

// Checks if the URL matches any given pattern
func MatchesAnyPattern(rawURL string, patterns []string) bool {
	for _, pattern := range patterns {
		matched, _ := regexp.MatchString(pattern, rawURL)
		if matched {
			return true
		}
	}
	return false
}
