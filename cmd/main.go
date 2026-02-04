package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"math-skills/stats"
)

// Program entrypoint and argument handling
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <data_file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	numbers, err := readNumbersFromFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Compute statistical metrics
	avg := stats.Average(numbers)
	med := stats.Median(numbers)
	varianceVal := stats.Variance(numbers)
	stdDevVal := stats.StandardDeviation(numbers)

	// Output rounded results
	fmt.Println("Average:", int(math.Round(avg)))
	fmt.Println("Median:", int(math.Round(med)))
	fmt.Println("Variance:", int(math.Round(varianceVal)))
	fmt.Println("Standard Deviation:", int(math.Round(stdDevVal)))
}

// Load numeric values from file
func readNumbersFromFile(path string) ([]float64, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Parse file lines into float slice
	lines := strings.Split(string(content), "\n")
	var nums []float64

	// Convert each line to a number
	for _, line := range lines {
		if line == "" {
			continue
		}
		value, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("invalid number %q: %w", line, err)
		}
		nums = append(nums, float64(value))
	}

	return nums, nil
}
