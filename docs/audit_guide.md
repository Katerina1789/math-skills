# Audit Guide

## Functional

### Step 1: Download the Test File
Download the auditor [test file](stat-bin-dockerized.zip).

### Step 2: Run the Script
After downloading the file, run the script:
```bash
./bin/math-skills
# or
./run.sh math-skills
```

This creates a `data.txt` file with random numbers.

### Step 3: Run Student Program
Run the student program with the created file:
```bash
go run cmd/main.go data.txt
```

### Step 4: Compare Outputs
Are the outputs of both programs (the one provided and the student one) in the same format?

In the output of the student program, are the data types of the values rounded integers?

Did the values of both programs match?

### Step 5: Repeat Testing
Do the same procedure (running the script provided and the student program) 3 more times in order to test new data sets.

Did the values of both programs match in all tries?

## Expected Output Format
```
Average: <integer>
Median: <integer>
Variance: <integer>
Standard Deviation: <integer>
```

## Bonus
Did the student provide a README with an explanation on how to test his/her program?
