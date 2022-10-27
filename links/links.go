package links

import (
	"bufio"
	"log"
	"net/url"
	"os"
	"strings"
)

type LinkStats struct {
	Internal int
	External int
	Error    int
}

// CountLinks evaluate passed links and identify links pointing to internal and external resources.
func CountLinks(rootDomain string, links []string) LinkStats {
	var root = strings.ToLower(rootDomain)

	var stats = LinkStats{
		Internal: 0,
		External: 0,
		Error:    0,
	}

	for _, link := range links {
		parsedURL, err := url.Parse(link)

		if err != nil {
			stats.Error++
			continue
		}

		host := strings.ToLower(parsedURL.Host)

		if host == root || host == "www."+root { // direct link to root domain
			stats.Internal++
			continue
		}

		if strings.HasPrefix(host, ".") { // relative path
			stats.Internal++
			continue
		}

		if parsedURL.Scheme == "" { // reference to file in the same directory
			stats.Internal++
			continue
		}

		stats.External++
	}

	return stats
}

// ReadLinksFromFile read links from file line by line
func ReadLinksFromFile(file string) ([]string, error) {
	readFile, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		_, err := url.Parse(line)
		if err != nil {
			log.Printf("Failed to parse url %s", line)
			continue
		}

		fileLines = append(fileLines, line)
	}

	return fileLines, nil
}
