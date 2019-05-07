// Sample program to show how to use a mutex to define critical
// sections of code that needs synchronous access.

package main

import (
	"fmt"
	"sync"
)

// counter is a variable incremented by all goroutines.
var counter int

// mutex is use to define a critical section of code.
var mutex sync.Mutex

func main() {

	// Number of goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two groutines.
	for i := 0; i < grs; i++ {
		go func() {
			for i := 0; i < 2; i++ {

				// Only allow one goroutine through this critical section at a time.
				mutex.Lock()
				{
					// Capture the value of counter.
					value := counter

					// Increment our local value of counter.
					value++

					// Store the value back into counter.
					counter = value
				}
				// Release the lock and allow any waiting goroutine through.
				mutex.Unlock()
			}
			wg.Done()
		}()
	}

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}
