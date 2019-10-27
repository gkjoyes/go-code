// Sample program to show how to declare, initialize and iterate over a map.
// Shows how iterating over a map is random.

package main

import "fmt"

// user represents someone using the program.
type user struct {
	name    string
	surname string
}

func main() {

	// Declare and initialize the map with values.
	users := map[string]user{
		"Roy":     user{name: "Rob", surname: "Roy"},
		"Ford":    user{name: "Henry", surname: "Ford"},
		"Mouse":   user{name: "Mickey", surname: "Mouse"},
		"Jackson": user{name: "Michael", surname: "Jackson"},
	}

	// Iterate over the map printing each key and value.
	for key, value := range users {
		fmt.Println(key, value)
	}

	fmt.Println()

	// Iterate over the map printing just the keys.
	// Notice the results are different.
	for key := range users {
		fmt.Println(key)
	}
}
