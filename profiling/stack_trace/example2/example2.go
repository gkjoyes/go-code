// Sample program to show how to read a stack trace when it packs values.
package main

func main() {
	example(true, false, true, 25)
}

//go:noinline
func example(b1, b2, b3 bool, i uint8) {
	panic("Want stack trace")
}

/*
	panic: Want stack trace

	goroutine 1 [running]:
	main.example(0xc019010001)
		/stack_trace/example2/example2.go:10 +0x39
	main.main()
		/stack_trace/example2/example2.go:5 +0x29
	exit status 2

	-> Declaration
	main.example(b1, b2, b3 bool, i uint8)

	-> Call
	main.example(true, false, true, 25)

	-> Stack trace
	main.example(0xc019010001)

	-> Word value (0xc019010001)
	Bits    Binary      Hex   Value
	00-07   0000 0001   01    true
	08-15   0000 0000   00    false
	16-23   0000 0001   01    true
	24-31   0001 1001   19    25
*/
