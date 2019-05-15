// This sample program demonstrates how to use the pool package to use a pool
// of goroutines to get work done.

package main

import (
	"log"
	"sync"
	"time"

	"github.com/george-kj/go-code/concurrency/patterns/pool"
)

// names provides a set of names to display.
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter provides special support for printing names.
type namePrinter struct {
	name string
}

// Work implements the Worker interface.
func (m namePrinter) Work() {
	log.Println(m.name)
	time.Sleep(3 * time.Second)
}

func main() {
	const routines = 10

	// Create a task pool.
	t := pool.New(routines)

	var wg sync.WaitGroup
	wg.Add(routines * len(names))

	for i := 0; i < routines; i++ {

		// Iterate over the slice of names.
		for _, name := range names {

			// Create a namePrinter and provide the specific name.
			np := namePrinter{
				name: name,
			}

			go func(np namePrinter) {

				// Submit the task to be worked on. When Do returns, we know it is
				// being handled.
				t.Do(np)
				wg.Done()
			}(np)
		}
	}
	wg.Wait()

	// Shutdown the task pool and wait for all existing work to be completed.
	t.Shutdown()
}
