package core

import (
	"fmt"
	"github.com/denis256/link-counter/links"
	"github.com/denis256/link-counter/page"
	"log"
	u "net/url"
	"strings"
)

type LinkCounter struct {
	ConcurrentRequests int

	Results []ScanResult
}

type ScanResult struct {
	PageUrl  string
	Internal int
	External int
	Success  bool
	Error    string
}

func scanUrl(url string) ScanResult {
	result := ScanResult{
		PageUrl: url,
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

	if strings.HasPrefix(root, "www.") {
		root = strings.TrimPrefix(root, "www.")
	}

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
