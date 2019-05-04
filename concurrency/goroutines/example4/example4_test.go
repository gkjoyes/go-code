package main

import (
	"math"
	"runtime"
	"testing"
)

func BenchmarkSingle(b *testing.B) {

	// Clear everything before start.
	runtime.GC()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		single(n)
	}
}

func BenchmarkUnlimited(b *testing.B) {

	// Clear everything before start.
	runtime.GC()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		unlimited(n)
	}
}

func BenchmarkNumCPU(b *testing.B) {

	// Clear everything before start.
	runtime.GC()
	b.ResetTimer()

	// Find out how deep we can create goroutine.
	maxLevel := int(math.Log2(float64(runtime.NumCPU())))

	for i := 0; i < b.N; i++ {
		numCPU(n, 0, maxLevel)
	}
}
