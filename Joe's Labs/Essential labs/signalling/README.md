# Signalling Exercise

## Overview

This exercise demonstrates basic synchronisation between two goroutines using an unbuffered channel. One goroutine signals that **Part A** is complete, while the second goroutine waits for that signal before continuing.

## Requirements

- Go **1.27.1** or later, matching the version declared in the parent `go.mod`.

## How to Run

Open a terminal in the `Essential labs` directory and run:

```bash
go run ./signalling
```

Alternatively, open a terminal in this folder and run:

```bash
go run .
```

## Expected Output

The output should follow this sequence:

```text
StuffOne - Part A
StuffTwo - Part A
StuffTwo - PartB
StuffOne - PartB
```

The second goroutine waits for five seconds before sending its message and receiving the signal.

## Implementation Notes

- Two goroutines perform separate tasks.
- An unbuffered channel passes a synchronisation signal between them.
- A `sync.WaitGroup` waits for both goroutines to finish.
