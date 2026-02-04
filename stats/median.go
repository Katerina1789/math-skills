package stats

import "sort"

// Median finds the middle value when numbers are sorted
func Median(nums []float64) float64 {
	cpy := make([]float64, len(nums))
	copy(cpy, nums)
	sort.Float64s(cpy)

	n := len(cpy)
	mid := n / 2

	if n%2 == 1 {
		return cpy[mid]
	}
	return (cpy[mid-1] + cpy[mid]) / 2
}
