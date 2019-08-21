package bounds

import "testing"

var sum int

func BenchmarkForward(b *testing.B) {
	ints := []int{}
	for i := 0; i < 5; i++ {
		ints = append(ints, i)
	}
	var s int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s = forward(ints)
	}
	sum = s
}

func BenchmarkBackwards(b *testing.B) {
	ints := []int{}
	for i := 0; i < 5; i++ {
		ints = append(ints, i)
	}
	var s int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s = backwards(ints)
	}
	sum = s
}

func BenchmarkRange(b *testing.B) {
	ints := []int{}
	for i := 0; i < 5; i++ {
		ints = append(ints, i)
	}
	var s int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s = rangeBounds(ints)
	}
	sum = s
}

/*
// section: output
BenchmarkForward-8      2000000000               0.94 ns/op
BenchmarkBackwards-8    2000000000               0.41 ns/op
// section: output
*/
