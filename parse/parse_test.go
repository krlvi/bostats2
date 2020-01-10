package parse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1pageBelowLimit(t *testing.T) {
	pages, err := totalPages("1 - 12", " av 12")
	assert.Nil(t, err)
	assert.Equal(t, 1, pages)
}

func Test_1pageAtLimit(t *testing.T) {
	pages, err := totalPages("1 - 50", " av 50")
	assert.Nil(t, err)
	assert.Equal(t, 1, pages)
}

func Test_2pages(t *testing.T) {
	pages, err := totalPages("1 - 50", " av 96")
	assert.Nil(t, err)
	assert.Equal(t, 2, pages)
}

func Test_9pages(t *testing.T) {
	pages, err := totalPages("1 - 50", " av 404")
	assert.Nil(t, err)
	assert.Equal(t, 9, pages)
}
