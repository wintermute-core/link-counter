package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountLinksSingle(t *testing.T) {

	stats := CountLinks("abc.com",
		[]string{
			"http://xyz.com",
			"http://abc.com",
			"http://111.com",
		})

	assert.Equal(t, 1, stats.Internal)
	assert.Equal(t, 2, stats.External)
	assert.Equal(t, 0, stats.Error)

}

func TestCountLinksDifferentSchemas(t *testing.T) {

	stats := CountLinks("abc.com",
		[]string{
			"http://abc.com",
			"https://abc.com",
			"/",
			"http://1.com",
			"https://2.com",
			"https://2.com/",
		})

	assert.Equal(t, 3, stats.Internal)
	assert.Equal(t, 3, stats.External)
	assert.Equal(t, 0, stats.Error)

}
