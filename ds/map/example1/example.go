// Sample program to show how to initialize a map, write to it,
// then read and delete from it.

package main

import (
	"fmt"
)

// user represents someone using the program.
type user struct {
	name    string
	surname string
}

func main() {

	// Declare and make a map that stores values of type user with
	// a key of type string.
	users := make(map[string]user)

	// Add key/value pairs to the map.
	users["Roy"] = user{name: "Roy", surname: "Roy"}
	users["Ford"] = user{name: "Henry", surname: "Ford"}
	users["Mouse"] = user{name: "Mickey", surname: "Mouse"}
	users["Jackson"] = user{name: "Michael", surname: "Jackson"}

	// Read the value at a specific key.
	mouse := users["Mouse"]
	fmt.Printf("%+v\n", mouse)

	// Replace the value at the Mouse key.
	users["Mouse"] = user{name: "jerry", surname: "Mouse"}

	// Read the Mouse key again.
	fmt.Printf("%+v\n", users["Mouse"])

	// Delete the value at a specific key.
	delete(users, "Roy")

	// Check the length of the map. There are only 3 elements.
	fmt.Println(len(users))

	// It is safe to delete an absent key.
	delete(users, "Roy")
}
