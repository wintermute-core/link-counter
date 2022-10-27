package core

import (
	"fmt"
	"log"
	u "net/url"
	"strings"
	"sync"

	"github.com/denis256/link-counter/links"
	"github.com/denis256/link-counter/page"
)

// LinkCounter structure to configure link counter application
type LinkCounter struct {
	Workers int
}

// ScanResult structure to store scan results of single url.
type ScanResult struct {
	PageURL  string
	Internal int
	External int
	Success  bool
	Error    string
}

// Scan passed url for internal and external links.
func (counter LinkCounter) Scan(urls []string) []ScanResult {
	var wg sync.WaitGroup

	jobs := make(chan string)
	results := make(chan ScanResult, len(urls))

	// start workers
	for i := 0; i < counter.Workers; i++ {
		wg.Add(1)

		go worker(&wg, jobs, results)
	}

	// sending each url for processing
	for _, u := range urls {
		jobs <- u
	}

	// closing job channel and wait for results
	close(jobs)
	wg.Wait()
	close(results)

	var result = make([]ScanResult, len(urls))

	var id = 0

	for r := range results {
		result[id] = r
		id++
	}

	return result
}

// worker function that process urls from incoming channel and send result to output channel
func worker(wg *sync.WaitGroup, jobs <-chan string, result chan ScanResult) {
	defer wg.Done()

	for job := range jobs {
		result <- scanURL(job)
	}
}

// scanURL download and scan urls from a single page.
func scanURL(url string) ScanResult {
	result := ScanResult{
		PageURL: url,
		Success: true,
	}

	// parse url and extract hostname
	parsedURL, err := u.Parse(url)

	if err != nil {
		result.Success = false
		result.Error = fmt.Sprint("%w", err)

		return result
	}

	root := strings.ToLower(parsedURL.Host)
	root = strings.TrimPrefix(root, "www.")

	// download page and process links
	log.Printf("Processing %s", url)
	pageContent, err := page.DownloadPage(url)

	if err != nil {
		result.Success = false
		result.Error = fmt.Sprint("%w", err)

		return result
	}

	parsedLinks := page.ParsePageLinks(&pageContent)
	stats := links.CountLinks(root, parsedLinks)

	// printing in logs that not all lings got parsed from page
	if stats.Error != 0 {
		log.Printf("Failed to parse %d links on page %s", stats.Error, url)
	}

	result.External = stats.External
	result.Internal = stats.Internal

	return result
}
