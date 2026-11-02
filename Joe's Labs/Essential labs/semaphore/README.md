# Semaphore Exercise

## Overview

This exercise demonstrates how a buffered channel can be used as a semaphore to control concurrency. The program launches 20 tasks but allows a maximum of five tasks to run at the same time.

## Requirements

- Go **1.27.1** or later, matching the version declared in the parent `go.mod`.

## How to Run

Open a terminal in the `Essential labs` directory and run:

```bash
go run ./semaphore
```

Alternatively, open a terminal in this folder and run:

```bash
go run .
```

## Expected Behaviour

The program prints a message for each of the 20 tasks. Each task simulates work by sleeping for two seconds. No more than five tasks should hold a semaphore slot at the same time.

The order of the task messages is not guaranteed and may differ between runs.

## Implementation Notes

- A buffered channel with a capacity of five acts as the semaphore.
- Sending to the channel acquires a slot.
- Receiving from the channel releases a slot.
- A `sync.WaitGroup` waits for all 20 tasks to finish.
