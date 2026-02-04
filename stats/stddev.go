package stats

import "math"

// StandardDeviation calculates the square root of variance
func StandardDeviation(nums []float64) float64 {
	return math.Sqrt(Variance(nums))
}
