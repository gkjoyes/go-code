package main

import (
	"fmt"
	"testing"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// TestLatency runs a single stream so we can look at blocking profiles
// for different buffer sizes.
func TestLatency(t *testing.T) {
	bufSize := 100

	fmt.Println("BufSize:", bufSize)
	stream(bufSize)
}

// TestLatencies provides a test to profile and trace channel latencies with
// a little data science sprinkled in.
func TestLatencies(t *testing.T) {
	var (
		bufSize int
		count   int
		first   time.Duration
	)

	pts := make(plotter.XYs, 20)

	for {

		// Perform a stream with specified buffer size.
		since := stream(bufSize)

		// Calculate how long this took and the percent of different from the
		// unbuffered channel.
		if bufSize == 0 {
			first = since
		}

		fmt.Printf("First: %v, Since: %v\n", first, since)
		dec := ((float64(first) - float64(since)) / float64(first)) * 100

		// Display the results.
		fmt.Printf("BufSize: %d\t%v\t%.2f%%\n", bufSize, since, dec)

		// Prepare the results for plotting.
		pts[count].X = float64(bufSize)
		pts[count].Y = dec
		count++

		// Want to look at a single buffer increment.
		if bufSize < 10 {
			bufSize++
			continue
		}

		// Increment by 10 moving forward.
		if bufSize == 100 {
			break
		}
		bufSize += 10
	}

}

// makePlot creates and saves a plot of the overall latencies
// difference from the unbuffered channel.
func makePlot(xys plotter.XYs) error {

	// Create a new plot.
	p, err := plot.New()
	if err != nil {
		return err
	}

	// Label the new plot.
	p.Title.Text = "Latencies (difference from the unbuffered channel)"
	p.X.Label.Text = "Buffer Length"
	p.Y.Label.Text = "Latency"

	// Add prepared points to the plot.
	if err = plotutil.AddLinePoints(p, "Latencies", xys); err != nil {
		return err
	}

	// Save the plot to a PNG file.
	return p.Save(10*vg.Inch, 5*vg.Inch, "latencies.png")
}
