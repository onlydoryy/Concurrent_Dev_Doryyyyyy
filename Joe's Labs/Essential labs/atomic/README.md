# Atomic Operations Exercise

## Overview

This exercise demonstrates how to safely update a shared value from multiple goroutines using Go's `sync/atomic` package. Ten goroutines each increment a shared atomic counter 1,000 times.

## Requirements

- Go **1.27.1** or later, matching the version declared in the parent `go.mod`.

## How to Run

Open a terminal in the `Essential labs` directory and run:

```bash
go run ./atomic
```

Alternatively, open a terminal in this folder and run:

```bash
go run .
```

## Expected Output

The program starts ten goroutines and should finish by printing:

```text
10000
```

The goroutine start messages may appear in a different order between runs.

## Implementation Notes

- `atomic.Int64` stores the shared counter.
- The `Add` method performs each increment atomically.
- A `sync.WaitGroup` waits for all ten goroutines to finish.
