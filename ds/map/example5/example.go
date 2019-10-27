// Sample program to show how to walk through a map by alphabetic key order.

package main

import (
	"fmt"
	"sort"
)

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

	// Pull the keys from the map.
	var keys []string
	for key := range users {
		keys = append(keys, key)
	}

	// Sort the keys alphabetically.
	sort.Strings(keys)

	// Walk through the keys and pull each value from the map.
	for _, k := range keys {
		fmt.Println(k, users[k])
	}
}
