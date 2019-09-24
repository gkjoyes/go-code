// Sample program to show how to read a stack trace.
package main

func main() {
	example(make([]string, 2, 4), "hello", 10)
}

//go:noinline
func example(slice []string, str string, i int) {
	panic("Want stack trace")
}

/*
	panic: Want stack trace

	goroutine 1 [running]:
	main.example(0xc000038710, 0x2, 0x4, 0x47345d, 0x5, 0xa)
		/stack_trace/example1/example1.go:10 +0x39
	main.main()
		/stack_trace/example1/example1.go:5 +0x72
	exit status 2

	-> Declaration
	main.example(slice []string, str string, i int)

	-> Call
	main.example(0xc000038710, 0x2, 0x4, 0x47345d, 0x5, 0xa)

	-> Stack trace
	main.example(0xc000038710, 0x2, 0x4, 0x47345d, 0x5, 0xa)

	-> Values
	Slice Value: 0xc000038710, 0x2, 0x4
	String Value: 0x47345d, 0x5
	Integer Value: 0xa

	Use `go build -gcflags -S` to map the PC offset values, +0x39 and +0x29 for each function call.
*/
