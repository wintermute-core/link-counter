package links

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountLinksSingle(t *testing.T) {
	stats := CountLinks("abc.com",
		[]string{
			"http://xyz.com",
			"http://abc.com",
			"http://abc.com/about.html",
			"http://111.com",
		})

	validateStats(t, stats, 2, 2, 0)
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

	validateStats(t, stats, 3, 3, 0)
}

func TestCountLinksRelativePath(t *testing.T) {
	stats := CountLinks("abc.com",
		[]string{
			"/",
			"..",
			".",
			"qwe.html",
		})

	validateStats(t, stats, 4, 0, 0)
}

func TestCountLinksErrors(t *testing.T) {
	stats := CountLinks("abc.com",
		[]string{
			"http//abc.com",
			"🐛U+1F41B\n",
			"http://abc.com/a.html",
			"http://xyz.com",
			"http://xyz.com/🐛.txt\t",
		})

	validateStats(t, stats, 2, 1, 2)
}

func TestReadLinksFromFile(t *testing.T) {
	links, err := ReadLinksFromFile("test_file_with_links.txt")
	assert.NoError(t, err)
	assert.NotEmpty(t, links)
	assert.Len(t, links, 2)
	assert.Contains(t, links, "http://google.com")
	assert.Contains(t, links, "http://qwe.com")
}

func TestReadLinksFromFileError(t *testing.T) {
	_, err := ReadLinksFromFile("not-existing-file-somewhere🪲.txt")
	assert.Error(t, err)
}

func validateStats(t *testing.T, stats LinkStats, internal, external, error int) {
	assert.Equalf(t, internal, stats.Internal, "Expected to have %d internal links", internal)
	assert.Equalf(t, external, stats.External, "Expected to have %d external links", external)
	assert.Equalf(t, error, stats.Error, "Expected to have %d errors", error)
}
