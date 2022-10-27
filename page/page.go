package page

import (
	"io"
	"net/http"
)

// DownloadPage - download url as string
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
