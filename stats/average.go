package stats

// Average calculates the mean of all numbers
func Average(nums []float64) float64 {
	sum := 0.0
	for _, v := range nums {
		sum += v
	}
	return sum / float64(len(nums))
}
