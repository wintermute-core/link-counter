package page

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBaseCloningSuite(t *testing.T) {
	content, err := DownloadPage("http://universal-development.com/contacts/")

	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(content))

}

func TestParsePageLinks(t *testing.T) {
	content, err := os.ReadFile("test_page.html")
	assert.NoError(t, err)

	links := ParsePageLinks(&content)
	assert.NotEqual(t, 0, len(links))

	assert.Contains(t, links, "/")
	assert.Contains(t, links, "..")
	assert.Contains(t, links, "https://github.com/denis256")
	assert.Contains(t, links, "https://universal-development.com/contacts/")
}
