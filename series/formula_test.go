package series

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetScgSeriesItemFunction(t *testing.T) {
	series := []int{3, 5, 9, 15, 23, 33, 45}
	for index, expect := range series {
		assert.Equal(t, expect, getScgSeriesItem(index))
	}
}

func BenchmarkGetScgSeriesItemFunction(b *testing.B) {
	for index := 0; index < b.N; index++ {
		getScgSeriesItem(index)
	}
}
