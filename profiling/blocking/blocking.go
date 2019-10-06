// Sample program to see what a trace will look like for basic channel latencies.
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

// data represents a set of bytes to process.
var data []byte

// init creates a data for processing.
func init() {
	f, err := os.Open("data.bytes")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err = ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bytes", len(data))
}

func main() {
	time := stream(10)
	fmt.Println(time)
}

// stream performs the moving of the data stream from one goroutine to the other.
func stream(bufSize int) time.Duration {

	// Create WaitGroup and channel.
	var wg sync.WaitGroup
	ch := make(chan int, bufSize)

	// Capture the reader for the input data.
	data := input()

	// Create the receiver goroutine.
	wg.Add(1)
	go func() {
		defer wg.Done()
		recv(ch)
	}()

	// Start the clock.
	st := time.Now()

	// Send all the data to the receiving goroutine.
	send(data, ch)

	// Close the channel and wait for the receiving goroutine to terminate.
	close(ch)
	wg.Wait()

	// Calculate how long the send took
	return time.Since(st)
}

// input returns a reader that can be used to read a stream of bytes.
func input() io.Reader {
	return bytes.NewBuffer(data)
}

// recv waits for bytes and add them up.
func recv(ch chan int) {
	var total int

	for v := range ch {
		total += v
	}
}

// send reads the bytes and sends them through the channel.
func send(r io.Reader, ch chan int) {
	buf := make([]byte, 1)

	for {
		n, err := r.Read(buf)
		if n == 0 || err != nil {
			break
		}

		ch <- int(buf[0])
	}
}
