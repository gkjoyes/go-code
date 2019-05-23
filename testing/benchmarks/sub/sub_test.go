// go test -run none -bench . -benchtime 3s -benchmem
// go test -run none -bench BenchmarkSprint/none -benchtime 3s -benchmem
// go test -run none -bench BenchmarkSprint/format -benchtime 3s -benchmem

// basic sub-benchmark test.
package sub_test

import (
	"fmt"
	"testing"
)

var gs string

// BenchmarkSprint tests all the Sprint related benchmarks as sub benchmarks.
func BenchmarkSprint(b *testing.B) {
	b.Run("none", benchmarkSprint)
	b.Run("format", benchmarkSprintf)
}

// benchmarkSprint tests the performance of using Sprint.
func benchmarkSprint(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello")
	}

	gs = s
}

// benchmarkSprintf tests the performance of using Sprintf.
func benchmarkSprintf(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello")
	}

	gs = s
}
