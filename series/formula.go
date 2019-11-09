package series

import "math"

// getScgSeriesItem function will calcute item of series by index.
func getScgSeriesItem(index int) int {
	square := math.Pow(float64(index), 2)
	return int(square) + index + 3
}
