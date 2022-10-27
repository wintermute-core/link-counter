package page

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBaseCloningSuite(t *testing.T) {
	content, err := DownloadPage("https://google.com")

	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(content))

}
