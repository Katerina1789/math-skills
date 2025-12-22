// Package main calculates statistical measures from a data file
package main

import (
	"bufio"   // Provides buffered I/O for efficient file reading
	"fmt"     // Formats and prints output to console
	"log"     // Handles error logging and program termination
	"math"    // Provides mathematical functions like Round and Sqrt
	"os"      // Accesses command-line arguments and file operations
	"sort"    // Sorts slices for median calculation
	"strconv" // Converts strings to numbers
)

// main is the entry point that orchestrates reading data and calculating statistics
func main() {
	// Check if a file path argument was provided
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <data_file>")
	}

	// Get the file path from command-line arguments
	filePath := os.Args[1]
	// Read all numbers from the file into a slice
	numbers, err := readNumbersFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate all required statistical measures
	avg := average(numbers)
	med := median(numbers)
	varianceVal := variance(numbers)
	stdDevVal := standardDeviation(numbers)

	// Print results as rounded integers in the required format
	fmt.Println("Average:", int(math.Round(avg)))
	fmt.Println("Median:", int(math.Round(med)))
	fmt.Println("Variance:", int(math.Round(varianceVal)))
	fmt.Println("Standard Deviation:", int(math.Round(stdDevVal)))
}

// readNumbersFromFile reads integers from a file and returns them as float64 slice
func readNumbersFromFile(path string) ([]float64, error) {
	// Open the file at the given path
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// Ensure file is closed when function exits
	defer file.Close()

	// Initialize slice to store numbers
	var nums []float64
	// Create scanner to read file line by line
	scanner := bufio.NewScanner(file)

	// Read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines
		if line == "" {
			continue
		}
		// Convert string to integer
		value, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("invalid number %q: %w", line, err)
		}
		// Append as float64 for statistical calculations
		nums = append(nums, float64(value))
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nums, nil
}

// average calculates the mean of all numbers in the slice
func average(nums []float64) float64 {
	// Initialize sum to zero
	sum := 0.0
	// Add each number to the sum
	for _, v := range nums {
		sum += v
	}
	// Divide total sum by count of numbers
	return sum / float64(len(nums))
}

// median finds the middle value when numbers are sorted
func median(nums []float64) float64 {
	// Create a copy to avoid modifying the original slice
	cpy := make([]float64, len(nums))
	copy(cpy, nums)

	// Sort numbers in ascending order
	sort.Float64s(cpy)

	// Get the length and middle index
	n := len(cpy)
	mid := n / 2

	// For odd length, return the middle element
	if n%2 == 1 {
		return cpy[mid]
	}
	// For even length, return average of two middle elements
	return (cpy[mid-1] + cpy[mid]) / 2.0
}

// variance measures how spread out the numbers are from the mean
func variance(nums []float64) float64 {
	// Calculate the mean first
	avg := average(nums)
	// Initialize sum of squared differences
	sumSquares := 0.0
	// Calculate squared difference from mean for each number
	for _, v := range nums {
		diff := v - avg
		sumSquares += diff * diff
	}
	// Divide by N for population variance
	return sumSquares / float64(len(nums))
}

// standardDeviation calculates the square root of variance
func standardDeviation(nums []float64) float64 {
	// Return the square root of the variance
	return math.Sqrt(variance(nums))
}
