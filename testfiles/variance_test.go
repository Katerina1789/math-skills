package stats

import (
	"math"
	"testing"

	"math-skills/stats"
)

func TestVariance(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{"basic", []float64{189, 113, 121, 114, 145, 110}, 784.6667},
		{"uniform", []float64{5, 5, 5, 5}, 0.0},
		{"simple", []float64{2, 4, 6, 8}, 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stats.Variance(tt.input)
			if math.Abs(result-tt.expected) > 0.01 {
				t.Errorf("Variance(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
