// Sample program to show how to create goroutines and
// how the goroutine scheduler behaves with two context.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	// Allocate two logical processors for the scheduler to use.
	runtime.GOMAXPROCS(2)
}

func main() {
	// wg is use to wait for the program to finish.
	// Add a count two, one for each goroutine.

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Declare an anonymous function and create a goroutine.
	go func() {

		// Display the alphabet three times.
		for count := 0; count < 3; count++ {
			for r := 'a'; r < 'z'; r++ {
				fmt.Printf("%c", r)
			}
		}

		// Tell main we are done.
		wg.Done()
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {

		// Display the alphabet three times.
		for count := 0; count < 3; count++ {
			for r := 'A'; r < 'Z'; r++ {
				fmt.Printf("%c", r)
			}

		}

		// Tell main we are done.
		wg.Done()
	}()

	// Wait for the goroutine to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}
