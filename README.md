# Math Skills

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/) [![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](https://opensource.org/licenses/MIT) [![Tests](https://img.shields.io/badge/Tests-Passed-brightgreen?style=for-the-badge)](docs/audit.md) [![Zone01](https://img.shields.io/badge/Zone01-Athens-blue?style=for-the-badge)](https://zone01.gr/gr/)

A professional Go program that calculates statistical measures from numerical datasets. Built as part of the Zone01 Athens curriculum.

---

## Description

This program reads a dataset from a file and computes four fundamental statistical measures:
- **Average** (Mean)
- **Median** (Middle value)
- **Variance** (Data spread)
- **Standard Deviation** (Square root of variance)

All results are output as rounded integers following the project specifications.

---

## Features

**Core Functionality**
- Calculates Average, Median, Variance, and Standard Deviation
- Handles datasets of any size
- Robust error handling for invalid input
- Standard library only (no external dependencies)

**Professional Structure**
- Modular package design with separated concerns
- Comprehensive unit tests for all functions
- Clean architecture following Go best practices
- Makefile for easy build and test automation

**Accurate Calculations**
- Population variance (not sample variance)
- Proper median calculation for even/odd counts
- Standard mathematical rounding

---

## Repository Structure

```
math-skills/
├── cmd/           //Application entry point
├── stats/         //Statistical calculation functions
├── testfiles/     //Unit tests for all functions
├── examples/      //Sample data files
├── docs/          //Project documentation
├── standard/      //License and contributing guidelines
├── Makefile
├── go.mod
├── .gitignore
└── README.md
```

---

## How to Run

### Quick Start
```bash
# Build and run with example data
make run

# Or manually
go run cmd/main.go examples/data1.txt
```

### Build Binary
```bash
# Build executable
make build

# Run the binary
./math-skills examples/data1.txt
```

### Run Tests
```bash
# Run all tests
make test

# Run with verbose output
go test ./... -v
```

### Other Commands
```bash
make clean    # Remove built binaries
make fmt      # Format code
make vet      # Check code quality
make help     # Show all commands
```

---

## Requirements

- **Go**: Version 1.21 or higher
- **OS**: Linux, macOS, or Windows
- **Dependencies**: None (standard library only)

### Input File Format
- One integer per line
- Empty lines are ignored
- Example:
  ```
  189
  113
  121
  114
  145
  110
  ```

### Output Format
```
Average: 132
Median: 118
Variance: 785
Standard Deviation: 28
```

---

## Algorithm Glossary

### Average (Arithmetic Mean)
The sum of all values divided by the count of values.

**Formula**: `μ = (Σxᵢ) / n`

**Example**: `[10, 20, 30]` → `(10+20+30)/3 = 20`

---

### Median
The middle value when data is sorted. For even counts, the average of the two middle values.

**Formula**: 
- Odd count: `x[(n+1)/2]`
- Even count: `(x[n/2] + x[n/2+1]) / 2`

**Example**: 
- `[1, 2, 3]` → `2`
- `[1, 2, 3, 4]` → `(2+3)/2 = 2.5`

---

### Variance (Population)
Measures how spread out numbers are from the mean. Uses population variance (divides by N, not N-1).

**Formula**: `σ² = Σ(xᵢ - μ)² / n`

**Example**: `[2, 4, 6]` → Mean=4, Variance=`((2-4)²+(4-4)²+(6-4)²)/3 = 8/3 ≈ 2.67`

---

### Standard Deviation
The square root of variance. Represents average distance from the mean.

**Formula**: `σ = √(σ²)`

**Example**: If variance = 4, then standard deviation = 2

---

## License

This project is licensed under the [MIT License](standard/LICENSE).
