# Mutex Exercise

## Overview

This exercise demonstrates how a mutex protects shared data when multiple goroutines access the same variable. Ten goroutines each increment a shared counter 1,000 times.

## Requirements

- Go **1.27.1** or later, matching the version declared in the parent `go.mod`.

## How to Run

Open a terminal in the `Essential labs` directory and run:

```bash
go run ./mutex
```

Alternatively, open a terminal in this folder and run:

```bash
go run .
```

## Expected Output

The program prints the goroutine numbers and then the final counter value:

```text
10000
```

The order of the goroutine messages may vary between runs.

## Implementation Notes

- A shared `int64` variable stores the counter.
- A `sync.Mutex` ensures only one goroutine increments the counter at a time.
- A `sync.WaitGroup` waits for all ten goroutines to finish.
