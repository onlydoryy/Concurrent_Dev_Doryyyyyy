# Dining Philosophers Exercise

## Overview

This exercise implements the classic Dining Philosophers concurrency problem. Five philosophers alternate between thinking and eating, sharing five forks. Each philosopher must acquire two forks before eating.

The implementation avoids the circular-wait condition by having the final philosopher acquire the forks in the opposite order. Each philosopher completes five eating cycles before the program exits.

## Requirements

- Go **1.27.1** or later, matching the version declared in the parent `go.mod`.

## How to Run

Open a terminal in the `Essential labs` directory and run:

```bash
go run ./DiningPhilosophers
```

Alternatively, open a terminal in this folder and run:

```bash
go run .
```

## Expected Behaviour

The program prints messages showing each philosopher thinking and eating. The delays are random, so the order of the messages will differ between runs. Once all five philosophers have completed their five cycles, the program prints:

```text
All philosophers have finished eating.
```

## Implementation Notes

- Goroutines represent the five philosophers.
- Buffered channels represent the forks.
- A `sync.WaitGroup` waits for all philosophers to finish.
- Random delays are used to simulate thinking and eating.
