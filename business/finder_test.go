package business

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBusinessSearch(t *testing.T) {
	apiKey := ""
	results, err := NewFinder(apiKey).Search(With("Bangsue", "restaurant"))
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.NotEmpty(t, results)
}

func BenchmarkBusinessSearch(b *testing.B) {
	apiKey := ""
	finder := NewFinder(apiKey)

	for i := 0; i < b.N; i++ {
		finder.Search(With("Bangsue", "restaurant"))
	}
}
