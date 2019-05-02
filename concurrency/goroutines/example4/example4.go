// Sample program to show you need to validate your benchmark results.
package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

// n contains the data to sort.
var n []int

// Generate the numbers to sort.
func init() {
	for i := 10; i >= 0; i-- {
		n = append(n, i)
	}
}

func main() {

	// Calculate how many levels deep we can create goroutines.
	maxLevel := int(math.Log2(float64(runtime.NumCPU())))

	fmt.Println("Final Response: ", numCPU(n, 0, maxLevel))
}

// single uses a single goroutine to perform the merge sort.
func single(n []int) []int {

	// Once we have a list of one we can begin to merge values.
	if len(n) <= 1 {
		return n
	}

	// Split the list in half.
	i := len(n) / 2

	// Sort the left side.
	l := single(n[:i])

	// Sort the right side.
	r := single(n[i:])

	// Place things in order and merge ordered list.
	return merge(l, r)
}

// unlimited uses a goroutine for every split to perform the merge sort.
func unlimited(n []int) []int {

	// Once we have a list of one we can begin to merge values.
	if len(n) <= 1 {
		return n
	}

	// Split the list in half.
	i := len(n) / 2

	// Maintain the ordered left and right side list.
	var l, r []int

	// For each split we will have 2 goroutines.
	var wg sync.WaitGroup
	wg.Add(2)

	// Sort the left side concurrently.
	go func() {
		l = unlimited(n[:i])
		wg.Done()
	}()

	// Sort the right side concurrently.
	go func() {
		r = unlimited(n[i:])
		wg.Done()
	}()

	// Wait for the spliting to end.
	wg.Wait()

	// Place the things in order and merge ordered lists.
	return merge(l, r)
}

func numCPU(n []int, lvl, maxLevel int) []int {

	// Once we have a list of one we can begin to merge values.
	if len(n) <= 1 {
		return n
	}

	// Split the list in half.
	i := len(n) / 2

	// Maintain the ordered left and right side lists.
	var l, r []int

	// We don't need more goroutines than we have logical processors.
	if lvl <= maxLevel {
		lvl++

		// For each split we will have 2 goroutines.
		var wg sync.WaitGroup
		wg.Add(2)

		// Sort the left side concurrently.
		go func() {
			l = numCPU(n[:i], lvl, maxLevel)
			wg.Done()
		}()

		// Sort the right side concurrently.
		go func() {
			r = numCPU(n[i:], lvl, maxLevel)
			wg.Done()
		}()

		// Wait for the spliting to end.
		wg.Wait()

		// Place things in order and merge ordered list.
		return merge(l, r)
	}

	// Sort the left and right side on this goroutine.
	l = numCPU(n[:i], lvl, maxLevel)
	r = numCPU(n[i:], lvl, maxLevel)

	// Place things in order and merge ordered lists.
	return merge(l, r)
}

// merge performs the merging to the two lists in proper order.
func merge(l, r []int) []int {

	// Declare the sorted return list with the proper capacity.
	ret := make([]int, 0, len(r)+len(l))

	// Compare the number of items required.
	for {
		switch {
		case len(l) == 0:

			// We appended everything in the left list so now append
			// everything contained in the right and return.
			return append(ret, r...)

		case len(r) == 0:

			// We appended everything in the right list so now append
			// everything contained in the left and return.
			return append(ret, l...)

		case l[0] <= r[0]:

			// First value in the left list is smaller than the
			// first value in the right list so append the left value.
			ret = append(ret, l[0])

			// Slice that first value away.
			l = l[1:]

		case l[0] >= r[0]:

			// First value in the right list is smaller than the
			// first value in the left list so append the right value.
			ret = append(ret, r[0])

			// Slice that first value away.
			r = r[1:]
		}
	}
}
