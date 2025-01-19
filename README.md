# Go Race Condition with Mutex and WaitGroup

This repository demonstrates a race condition that can occur in Go programs when using mutexes and wait groups. The bug is subtle, and it may not always manifest itself, making it difficult to detect.

## Bug Description

The program uses a mutex to protect a critical section of code. The main goroutine acquires the mutex, launches a goroutine that's supposed to release the mutex later, and then waits for the goroutine to finish. However, if the main goroutine exits before the goroutine has a chance to release the mutex, the program will deadlock.

## Solution

The solution involves ensuring that the main goroutine waits for the goroutine to finish before exiting. This prevents the program from deadlocking by guaranteeing that the mutex is released before the main goroutine exits.

## How to Reproduce

1. Clone this repository.
2. Run the `bug.go` file using `go run bug.go`.  The output might not always show a deadlock.
3. Run the `bugSolution.go` file using `go run bugSolution.go`.  This will demonstrate the fixed code.