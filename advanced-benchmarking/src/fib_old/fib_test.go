package fib

import "testing"

// Fib computes the n'th number in the Fibonacci series.
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// section: benchmark
var result int

func BenchmarkFib(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = Fib(20) // run the Fib function b.N times
	}
	result = r
}

// section: benchmark

func TestFib(t *testing.T) {
	fibs := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}
	for n, want := range fibs {
		got := Fib(n)
		if want != got {
			t.Errorf("Fib(%d): want %d, got %d", n, want, got)
		}
	}
}
