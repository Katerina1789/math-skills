# Math Skills

A Go program that calculates statistical measures (Average, Median, Variance, and Standard Deviation) from a data file as part of the Zone01 Athens curriculum.

## Repository Structure

```
math-skills/
├── main.go      # Main program with statistical calculations
├── data.txt     # Sample data file for testing
├── README.md    # Project documentation
└── LICENSE      # MIT License
```

## How to Run

```bash
go run main.go data.txt
```

## Input Format

The data file should contain one integer per line:
```
189
113
121
114
145
110
```

## Output Format

The program outputs four statistical measures as rounded integers:
```
Average: 132
Median: 118
Variance: 785
Standard Deviation: 28
```

## Testing

### For Auditors

1. Download the necessary file from [here](https://github.com/01-edu/public/blob/master/subjects/math-skills/audit/README.md)
2. Run the script with `./bin/math-skills` or `./run.sh math-skills`, generate a `data.txt` file, copy it to this repository
3. Run this program with the generated file:
   ```bash
   go run main.go data.txt
   ```
4. Compare the outputs to verify correctness

### Manual Testing

To test with the provided data file:
```bash
go run main.go data.txt
```

To test with your own data file:
```bash
go run main.go your_data_file.txt
```

## Example

With the provided `data.txt` containing:
```
189
113
121
114
145
110
```

The output will be:
```
Average: 132
Median: 118
Variance: 785
Standard Deviation: 28
```

## Requirements

- Go 1.11 or higher
- No external dependencies (uses only standard library)

## Functional Requirements

This program implements the following statistical calculations:
- **Average**: Sum of all values divided by count
- **Median**: Middle value when data is sorted (or average of two middle values for even count)
- **Variance**: Population variance (sum of squared differences from mean, divided by N)
- **Standard Deviation**: Square root of variance

## License

[MIT LICENSE](LICENSE)
