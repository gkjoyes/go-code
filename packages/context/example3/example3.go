// Sample program to show how to use the WithDeadline function.
package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	UserID string
}

func main() {

	// Set a deadline.
	deadline := time.Now().Add(150 * time.Millisecond)

	// Create a context that is both manually cancellable and will signal a cancel
	// at the specified data/time.
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// Create a channel to receive a signal for work is done.
	ch := make(chan data, 1)

	// Ask the goroutine to do some work for us.
	go func() {

		// Simulate work.
		time.Sleep(200 * time.Millisecond)

		// Report the work is done.
		ch <- data{"123"}
	}()

	// Wait for the work to finish. If it takes too long move on.
	select {
	case d := <-ch:
		fmt.Println("work completed", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}
