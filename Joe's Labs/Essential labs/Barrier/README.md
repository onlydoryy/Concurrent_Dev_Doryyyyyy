# Barrier Exercise

## Overview

This exercise demonstrates how to coordinate multiple goroutines using a barrier. Ten goroutines complete **Part A** before any of them proceed to **Part B**. A mutex protects the shared counter, while a channel is used to release the waiting goroutines.

## Requirements

- Go **1.27.1** or later, matching the version declared in the parent `go.mod`.
- The Go module dependency `golang.org/x/sync`, which is declared in the parent module.

## How to Run

Open a terminal in the `Essential labs` directory and run:

```bash
go run ./Barrier
```

Alternatively, open a terminal in this folder and run:

```bash
go run .
```

If dependencies have not already been downloaded, run `go mod download` from the `Essential labs` directory first.

## Expected Behaviour

The program starts ten goroutines. The output should show all ten `Part A` messages before the `PartB` messages begin. The order of goroutine numbers may vary because goroutines execute concurrently.

## Implementation Notes

- A `sync.Mutex` protects access to the shared counter.
- A channel is used to coordinate the barrier.
- A `sync.WaitGroup` keeps the main function running until all goroutines finish.
