package stats

// Variance measures how spread out numbers are from the mean
func Variance(nums []float64) float64 {
	avg := Average(nums)
	sumSquares := 0.0
	for _, v := range nums {
		diff := v - avg
		sumSquares += diff * diff
	}
	return sumSquares / float64(len(nums))
}
