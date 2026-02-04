# Architecture

## Project Structure

```
math-skills/
├── cmd/
│   └── main.go
├── stats/
│   ├── average.go
│   ├── median.go
│   ├── variance.go
│   └── stddev.go
├── testfiles/
│   ├── average_test.go
│   ├── median_test.go
│   ├── variance_test.go
│   └── stddev_test.go
├── examples/
│   ├── data1.txt
│   └── data2.txt
├── docs/
│   ├── audit_guide.md
│   ├── architecture.md
│   └── project_flowchart.md
├── standard/
│   ├── LICENSE
│   └── CONTRIBUTING.md
├── Makefile
├── go.mod
├── .gitignore
└── README.md
```

## File Descriptions

### cmd/
**main.go** - Application entry point. Validates command-line arguments, reads and parses the input file, calls stats package functions, rounds results to integers, and prints formatted output.

### stats/
**average.go** - Contains the Average function that calculates the arithmetic mean by summing all values and dividing by count.

**median.go** - Contains the Median function that finds the middle value in a sorted dataset. Returns the middle element for odd counts, or the average of two middle elements for even counts.

**variance.go** - Contains the Variance function that calculates population variance by computing the average of squared differences from the mean.

**stddev.go** - Contains the StandardDeviation function that calculates the square root of variance.

### testfiles/
**average_test.go** - Unit tests for the Average function covering basic cases, single values, negative numbers, and mixed values.

**median_test.go** - Unit tests for the Median function covering even/odd counts, single values, and unsorted data.

**variance_test.go** - Unit tests for the Variance function covering basic cases, uniform data, and simple datasets.

**stddev_test.go** - Unit tests for the StandardDeviation function covering basic cases, uniform data, and simple datasets.

### examples/
**data1.txt** - Sample dataset with 6 integers (189, 113, 121, 114, 145, 110) for testing and demonstration.

**data2.txt** - Additional sample dataset with 5 integers (10, 20, 30, 40, 50) for testing and demonstration.

### docs/
**audit_guide.md** - Instructions for auditors to verify program correctness using the official test script.

**architecture.md** - This file. Documents the project structure and explains what each file does.

**project_flowchart.md** - Visual flowchart showing the complete program flow with all functions and their interactions.

### standard/
**LICENSE** - MIT License for the project.

**CONTRIBUTING.md** - Development guidelines including setup instructions, available commands, code style rules, and contribution workflow.

### Root Files
**Makefile** - Build automation with targets for build, run, test, clean, fmt, vet, install, and help.

**go.mod** - Go module definition specifying module name (math-skills) and Go version (1.21).

**.gitignore** - Git ignore rules for binaries, test files, IDE files, OS files, and generated data.

**README.md** - Main project documentation with description, features, structure, usage instructions, requirements, and algorithm explanations.
