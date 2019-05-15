// This sample program demonstrates how the logger package works.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/george-kj/go-code/concurrency/patterns/logger"
)

// device allow us to mock a device we write log to.
type device struct {
	off bool
}

func main() {

	// Number of goroutines that will be writing logs.
	const grs = 10

	// Create a logger value with a buffer of capacity for each
	// goroutines that will be logging.
	var d device
	l := logger.New(&d, grs)

	// Generate goroutines, each writing to disk.
	for i := 0; i < grs; i++ {
		go func(id int) {
			for {
				l.Write(fmt.Sprintf("%d: log data", id))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	// We want to control the simulated disk blocking. Capture interrupt
	// signals to toggle device issues. Use <ctrl> Z to kill the program.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	for {
		<-sigChan

		// Here we have data race. Let's keep things simple to show the mechanics.
		d.off = !d.off
	}

}

// Write implements the io.Writer interface.
func (d *device) Write(p []byte) (n int, err error) {
	if d.off {

		// Simulate disk problem.
		time.Sleep(time.Second)
	}
	fmt.Println(string(p))
	return len(p), nil
}
