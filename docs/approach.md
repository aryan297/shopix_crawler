# Approach to Solving the Problem

## Problem Definition
The goal of this project is to develop a scalable and efficient web crawler that can crawl multiple e-commerce websites and extract product URLs. The crawler needs to handle dynamic content, especially Single Page Applications (SPAs) that load product links using JavaScript.

## Steps to Solve the Problem

### 1. **Choosing the Right Technology**
We need a solution that:
- Can handle dynamic content.
- Is scalable and efficient for multiple domains.
- Can extract URLs that match specific patterns for product pages.

**Technology Chosen**:
- **Go (Golang)**: It is known for its performance, concurrency handling (goroutines), and suitability for scalable applications.
- **Chromedp**: A Go package that can control a headless Chrome browser, making it suitable for SPAs, allowing us to simulate interactions and scrape content that is rendered dynamically with JavaScript.
- **Utils Package**: Custom utilities to match product URLs using regular expressions for common product page URL patterns.

### 2. **Crawling Process**
The crawler follows the following steps for each e-commerce site:
1. **Initial Setup**:
   - Create a new **Chromedp** context with a 60-second timeout.
   
2. **Navigating to the Site**:
   - The crawler visits the base URL of the website using **chromedp.Navigate**.

3. **Waiting for the Page to Load**:
   - We ensure the page is fully loaded by waiting for the `<body>` tag to be available using **chromedp.WaitReady**.
   
4. **Extracting Links**:
   - Once the page is fully loaded, we extract all URLs using **chromedp.Evaluate** to run JavaScript that gathers all links from the `<a>` tags on the page.

5. **Filtering Product Links**:
   - After gathering all links, we filter them using predefined product URL patterns, such as `/product/`, `/item/`, and `/p/`. These patterns are stored in the **utils** package.
   
6. **Storing Results**:
   - The valid product URLs are stored in a unique set and saved as a JSON file in the `output/results.json` directory.

### 3. **Handling Dynamic Content (SPA)**
Many modern e-commerce websites use **Single Page Applications (SPAs)** that load their product data dynamically via JavaScript. Traditional scraping tools that rely on static HTML parsing wouldn't work for these sites.

To handle this, we:
- Use **Chromedp** to control a headless browser that can render JavaScript and interact with dynamic content.
- Extract URLs only after all JavaScript content is loaded, ensuring we capture all product URLs visible on the page.

### 4. **Concurrency and Scalability**
To scale this solution for large websites with many products:
- The current solution can be modified to use **goroutines** in Go, which would allow crawling multiple pages in parallel.
- This approach ensures the crawler remains efficient even when handling numerous domains.

### 5. **Handling Edge Cases**
- **Slow Page Loads**: We set a timeout to avoid waiting indefinitely for pages that are slow to load.
- **Invalid URLs**: We ensure that URLs are valid and match product patterns. If the page contains invalid or irrelevant URLs, they are filtered out.
- **Custom URL Structures**: Some websites may use custom URL structures for product pages. We can extend the `productPatterns` list to handle these variations.

### 6. **How the Crawler Handles Errors**
- **Navigation Errors**: If a site fails to load, the crawler logs an error and continues to the next domain.
- **Timeouts**: If the page takes too long to load, it times out and skips the page.

---

## Future Improvements:
- **Parallel Crawling**: Extend the crawler to handle concurrent or parallel crawling of multiple pages to improve performance for large sites.
- **Robust Pattern Matching**: Enhance the product URL patterns to match a wider variety of e-commerce site structures.
- **Error Handling**: Add retries or error recovery mechanisms to deal with temporary issues like slow server responses.