package stats

import (
	"math"
	"testing"

	"math-skills/stats"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{"basic", []float64{189, 113, 121, 114, 145, 110}, 132.0},
		{"single", []float64{50}, 50.0},
		{"negative", []float64{-10, -20, -30}, -20.0},
		{"mixed", []float64{10, -10, 20, -20}, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stats.Average(tt.input)
			if math.Abs(result-tt.expected) > 0.01 {
				t.Errorf("Average(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
