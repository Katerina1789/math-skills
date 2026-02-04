package stats

import (
	"math"
	"testing"

	"math-skills/stats"
)

func TestMedian(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{"even_count", []float64{189, 113, 121, 114, 145, 110}, 117.5},
		{"odd_count", []float64{1, 2, 3, 4, 5}, 3.0},
		{"single", []float64{42}, 42.0},
		{"unsorted", []float64{5, 1, 3, 2, 4}, 3.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stats.Median(tt.input)
			if math.Abs(result-tt.expected) > 0.01 {
				t.Errorf("Median(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
