package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsingLinks(t *testing.T) {

	result := scanUrl("https://universal-development.com/contacts/")

	assert.True(t, result.Success)
	assert.Equal(t, "https://universal-development.com/contacts/", result.PageUrl)
	assert.Equal(t, 5, result.Internal)
	assert.Equal(t, 7, result.External)

}
