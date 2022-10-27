package page

import (
	"bytes"
	"golang.org/x/net/html"
	"io"
	"net/http"
)

// DownloadPage download url as string.
func DownloadPage(url string) ([]byte, error) {

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body := response.Body
	defer body.Close()

	content, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// ParsePageLinks parse and extract links from page.
func ParsePageLinks(page *[]byte) []string {
	var links []string
	tokenizer := html.NewTokenizer(bytes.NewReader(*page))
	for {
		token := tokenizer.Next()
		switch {
		case token == html.ErrorToken:
			return links
		case token == html.StartTagToken:
			t := tokenizer.Token()
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}
			url, found := fetchHref(t)
			if !found {
				continue
			}
			links = append(links, url)
		}
	}
}

// fetchHref extract HREF value of a link.
func fetchHref(token html.Token) (string, bool) {
	found := false
	link := ""
	for _, a := range token.Attr {
		if a.Key == "href" {
			link = a.Val
			found = true
		}
	}
	return link, found
}
