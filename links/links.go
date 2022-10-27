package links

import (
	"net/url"
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
