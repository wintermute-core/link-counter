package core

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingLinks(t *testing.T) {
	result := scanURL("https://universal-development.com/contacts/")

	assert.True(t, result.Success)
	assert.Equal(t, "https://universal-development.com/contacts/", result.PageURL)
	assert.Equal(t, 5, result.Internal)
	assert.Equal(t, 7, result.External)
}

func TestParsingLinksWww(t *testing.T) {
	result := scanURL("https://www.universal-development.com/contacts/")

	assert.True(t, result.Success)
	assert.Equal(t, "https://www.universal-development.com/contacts/", result.PageURL)
	assert.Equal(t, 5, result.Internal)
	assert.Equal(t, 7, result.External)
}

func TestDownloadingBrokenUrl(t *testing.T) {
	result := scanURL("http://not-resolving-domain-name-for-sure-1111.com")

	assert.False(t, result.Success)
	assert.NotEmpty(t, result.Error)
}

func TestMultipleLinksEvaluation(t *testing.T) {
	counter := LinkCounter{Workers: 2}

	results := counter.Scan([]string{
		"https://universal-development.com/contacts/",
		"https://universal-development.com/",
		"http://universal-development.com/services/",
		"http://universal-development.com/products/",
	})

	assert.NotEmpty(t, results)
	assert.Equal(t, 4, len(results))
}

func TestSerialization(t *testing.T) {
	result := ScanResult{
		PageURL:  "http://google.com",
		Internal: 666,
		External: 111,
		Success:  true,
		Error:    "Test",
	}

	bytes, err := json.Marshal(result)
	assert.NoError(t, err)

	rawJSON := map[string]interface{}{}

	err = json.Unmarshal(bytes, &rawJSON)
	assert.NoError(t, err)

	assert.Equal(t, rawJSON["page_url"], result.PageURL)
	assert.Equal(t, fmt.Sprintf("%v", rawJSON["internal_links_num"]), fmt.Sprintf("%v", result.Internal))
	assert.Equal(t, fmt.Sprintf("%v", rawJSON["external_links_num"]), fmt.Sprintf("%v", result.External))
	assert.Equal(t, rawJSON["success"], result.Success)
	assert.Equal(t, rawJSON["error_message"], result.Error)
}
