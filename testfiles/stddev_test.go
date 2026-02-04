package stats

import (
	"math"
	"testing"

	"math-skills/stats"
)

func TestStandardDeviation(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{"basic", []float64{189, 113, 121, 114, 145, 110}, 28.01},
		{"uniform", []float64{10, 10, 10}, 0.0},
		{"simple", []float64{2, 4, 6, 8}, 2.236},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stats.StandardDeviation(tt.input)
			if math.Abs(result-tt.expected) > 0.01 {
				t.Errorf("StandardDeviation(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
