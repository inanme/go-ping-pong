package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	withCancel()
}

func withTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go doSomething(ctx, 6*time.Second)
	go iterateOverSomething(ctx, 8)

	fmt.Scanln()
}

func withCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 5)
		cancel()
	}()

	go doSomething(ctx, 6*time.Second)
	go iterateOverSomething(ctx, 8)

	fmt.Scanln()
}

func doSomething(ctx context.Context, timeSleep time.Duration) {
	fmt.Println("Starting doSomething")
	select {
	// Here I'm simulating the time doSomething will take to process.
	case <-time.After(timeSleep):
		fmt.Println("Finished doSomething")
	// Check the context Done to get cancellation signal
	case <-ctx.Done():
		fmt.Println("Cancelling doSomething")
	}
}

func iterateOverSomething(ctx context.Context, iterations int) {
	for i := 0; i < iterations; i++ {
		switch {
		case errors.Is(ctx.Err(), context.Canceled):
			fmt.Println("Leaving iterateOverSomething... Cancelled")
			return
		case errors.Is(ctx.Err(), context.DeadlineExceeded):
			fmt.Println("Leaving iterateOverSomething... DeadlineExceeded")
			return
		case ctx.Err() != nil:
			fmt.Println("Leaving iterateOverSomething... ", ctx.Err())
			return
		}
		time.Sleep(time.Second * 1)
		fmt.Println("Iteration ", i)
	}
}
