// Sample program to show how only types that can have equality defined on them
// can be a map key.

package main

// user represents someone using the program.
type user struct {
	name    string
	surname string
}

// users defines a set of users.
type users []user

func main() {

	// Declare and make a map that uses a slice as the key.
	// u := make(map[users]int)

	// ./example.go:17:12: invalid map key type users
}
