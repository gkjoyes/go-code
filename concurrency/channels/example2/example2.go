// Sample program to show how to use an unbuffered channel to simulate a game of
// tennis between two goroutines.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Create a unbuffered channel.
	court := make(chan int)

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(2)

	// Lunch two players.
	go func() {
		defer wg.Done()
		player("Serena", court)
	}()

	go func() {
		defer wg.Done()
		player("Venus", court)
	}()

	// Start the set.
	court <- 1

	wg.Wait()
}

// player simulates a person playing the game of tennis.
func player(name string, court chan int) {
	for {
		// Wait for the ball to hit back to us.
		ball, wd := <-court
		if !wd {
			// If channel was closed we won.
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Pick a random number and see if we miss the ball.
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// Close the channel to signal we lost.
			close(court)
			return
		}

		// Display and then increment the hit count by one.
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player.
		court <- ball
	}
}
