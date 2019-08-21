package mem_test

import (
	"strings"
	"testing"
)

var sink int

// section: benchmark
func BenchmarkIndex(b *testing.B) {
	s := "Simplicity is the ultimate sophistication."
	var i int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		i = strings.Index(s, "the")
	}
	sink = i
}

// section: benchmark

/*
// section: output-call
goos: darwin
goarch: amd64
pkg: github.com/gopherguides/learn/_training/advanced/benchmarking/src/benchmem
BenchmarkIndexCall-8    200000000                8.87 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/gopherguides/learn/_training/advanced/benchmarking/src/benchmem      2.680s
// section: output-call


*/
