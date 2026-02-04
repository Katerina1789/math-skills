# Program Flowchart

```
                                    START
                                      │
                                      ▼
                        ┌─────────────────────────────┐
                        │   main()                    │
                        │   Check command-line args   │
                        └─────────────┬───────────────┘
                                      │
                         ┌────────────┴────────────┐
                         │                         │
                    len(os.Args) < 2?         len(os.Args) >= 2
                         │                         │
                         ▼                         ▼
              ┌──────────────────────┐   ┌─────────────────────┐
              │ Print usage message  │   │filePath = os.Args[1]│
              │ os.Exit(1)           │   └──────────┬──────────┘
              └──────────────────────┘              │
                                                     ▼
                                      ┌──────────────────────────┐
                                      │ readNumbersFromFile()    │
                                      │ Read file content        │
                                      └──────────┬───────────────┘
                                                 │
                                    ┌────────────┴────────────┐
                                    │                         │
                               File error?                File OK
                                    │                         │
                                    ▼                         ▼
                         ┌──────────────────┐    ┌────────────────────────┐
                         │ Return error     │    │ Split content by lines │
                         │ Print error      │    └──────────┬─────────────┘
                         │ os.Exit(1)       │               │
                         └──────────────────┘               ▼
                                              ┌──────────────────────────┐
                                              │ Loop through each line   │
                                              └──────────┬───────────────┘
                                                         │
                                            ┌────────────┴────────────┐
                                            │                         │
                                       Empty line?              Has content
                                            │                         │
                                            ▼                         ▼
                                    ┌──────────────┐    ┌─────────────────────┐
                                    │ Skip/Continue│    │ strconv.Atoi(line)  │
                                    └──────────────┘    └──────────┬──────────┘
                                                                   │
                                                      ┌────────────┴────────────┐
                                                      │                         │
                                                 Invalid number?           Valid number
                                                      │                         │
                                                      ▼                         ▼
                                           ┌──────────────────┐    ┌────────────────────┐
                                           │ Return error     │    │ Append to nums[]   │
                                           │ Print error      │    │ as float64         │
                                           │ os.Exit(1)       │    └──────────┬─────────┘
                                           └──────────────────┘               │
                                                                              ▼
                                                                   ┌──────────────────┐
                                                                   │ Return nums[]    │
                                                                   └────────┬─────────┘
                                                                            │
                                                                            ▼
                                                              ┌──────────────────────────┐
                                                              │ Calculate Statistics     │
                                                              └──────────┬───────────────┘
                                                                         │
                        ┌────────────────────────────────────────────────┼────────────────────────────────────────────────┐
                        │                                                │                                                │
                        ▼                                                ▼                                                ▼
          ┌──────────────────────────┐                    ┌──────────────────────────┐                  ┌──────────────────────────┐
          │ stats.Average(numbers)   │                    │ stats.Median(numbers)    │                  │ stats.Variance(numbers)  │
          └──────────┬───────────────┘                    └──────────┬───────────────┘                  └──────────┬───────────────┘
                     │                                               │                                             │
                     ▼                                               ▼                                             ▼
          ┌──────────────────────┐                      ┌──────────────────────────┐                  ┌──────────────────────────┐
          │ sum = 0.0            │                      │ Copy array               │                  │ avg = Average(nums)      │
          │ for each number:     │                      │ sort.Float64s(copy)      │                  │ sumSquares = 0.0         │
          │   sum += number      │                      └──────────┬───────────────┘                  │ for each number:         │
          │ return sum / count   │                                 │                                  │   diff = num - avg       │
          └──────────┬───────────┘                      ┌───────────┴────────────┐                    │   sumSquares += diff²    │
                     │                                  │                        │                    │ return sumSquares / N    │
                     │                             Odd count?              Even count                 └──────────┬───────────────┘
                     │                                  │                        │                               │
                     │                                  ▼                        ▼                               │
                     │                      ┌────────────────────┐  ┌─────────────────────────┐                  │
                     │                      │ return cpy[mid]    │  │ return (cpy[mid-1] +    │                  │
                     │                      └────────────────────┘  │         cpy[mid]) / 2   │                  │
                     │                                              └─────────────────────────┘                  │
                     │                                                                                           │
                     └──────────────────────────────────────────┬────────────────────────────────────────────────┘
                                                                │
                                                                ▼
                                                   ┌──────────────────────────────────┐
                                                   │ stats.StandardDeviation(numbers) │
                                                   └──────────────┬───────────────────┘
                                                                  │
                                                                  ▼
                                                      ┌────────────────────────┐
                                                      │ variance = Variance()  │
                                                      │ return √variance       │
                                                      └──────────┬─────────────┘
                                                                 │
                                                                 ▼
                                                   ┌──────────────────────────┐
                                                   │ Round all results using  │
                                                   │ math.Round() and convert │
                                                   │ to int                   │
                                                   └──────────┬───────────────┘
                                                              │
                                                              ▼
                                                   ┌──────────────────────────┐
                                                   │ Print formatted output:  │
                                                   │ Average: X               │
                                                   │ Median: X                │
                                                   │ Variance: X              │
                                                   │ Standard Deviation: X    │
                                                   └──────────┬───────────────┘
                                                              │
                                                              ▼
                                                             END
```

## Function Explanations

**main()** - Entry point that orchestrates the entire program flow. Validates arguments, reads file, calculates statistics, and prints results.

**readNumbersFromFile(path)** - Reads file content, splits by newlines, converts each line to integer, returns slice of float64 values.

**stats.Average(nums)** - Sums all numbers and divides by count to get arithmetic mean.

**stats.Median(nums)** - Copies array, sorts it, returns middle value (or average of two middle values for even count).

**stats.Variance(nums)** - Calculates average first, then computes sum of squared differences from mean, divides by N (population variance).

**stats.StandardDeviation(nums)** - Calls Variance function and returns its square root.
