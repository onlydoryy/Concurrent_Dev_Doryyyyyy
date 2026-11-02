# Rendezvous Exercise

## Overview

This exercise explores the rendezvous synchronisation pattern, where concurrent tasks must reach a synchronisation point before proceeding. Five goroutines perform **Part A** and then wait on a channel before attempting **Part B**.

## Requirements

- Go **1.27.1** or later, matching the version declared in the parent `go.mod`.

## How to Run

Open a terminal in the `Essential labs` directory and run:

```bash
go run ./Randezvous
```

Alternatively, open a terminal in this folder and run:

```bash
go run .
```

## Current Implementation Status

**The current source code does not complete successfully.** Each goroutine sends a signal on `aArrived` and then waits to receive from `bArrived`. The main function receives the arrival signals but does not send any values on `bArrived`. As a result, the goroutines remain blocked and the program eventually reports a deadlock.

This note documents the behaviour of the code currently in the repository; the rendezvous signalling still needs to be completed.

## Implementation Notes

- Goroutines represent the concurrent tasks.
- Unbuffered channels are used for synchronisation.
- A `sync.WaitGroup` is intended to wait for all goroutines to finish.
