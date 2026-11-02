# Semaphore and Collatz Exercise

## Overview

This exercise demonstrates how a weighted semaphore can limit the number of goroutines performing work concurrently. The program calculates the number of steps required for values 1 to 64 to reach 1 under the Collatz sequence.

## Requirements

- Go **1.27.1** or later, matching the version declared in the parent `go.mod`.
- The Go module dependency `golang.org/x/sync`, which is declared in the parent module.

## How to Run

Open a terminal in the `Essential labs` directory and run:

```bash
go run ./sem-ex
```

Alternatively, open a terminal in this folder and run:

```bash
go run .
```

If dependencies have not already been downloaded, run `go mod download` from the `Essential labs` directory first.

## Expected Behaviour

The program prints an integer slice containing the Collatz step count for each input from 1 to 64. The results are stored by input index, so the final slice is in numerical input order.

## Implementation Notes

- A weighted semaphore from `golang.org/x/sync/semaphore` limits concurrent work.
- Goroutines calculate the Collatz step count for individual values.
- The semaphore is used to wait for all outstanding calculations to finish.
