package main

import "testing"

// section: benchmark
var beatles = map[string]string{
	"John":   "guitar",
	"Paul":   "bass",
	"George": "guitar",
	"Ringo":  "drums",
}

var global string // prevent benchmark inlining

func BenchmarkConversion(b *testing.B) {
	var key = []byte{'J', 'o', 'h', 'n'}
	var r string
	for n := 0; n < b.N; n++ {
		r = beatles[string(key)]
	}
	global = r
}

func BenchmarkAllocation(b *testing.B) {
	var key = []byte{'J', 'o', 'h', 'n'}
	var r string
	for n := 0; n < b.N; n++ {
		k := string(key)
		r = beatles[k]
	}
	global = r
}

// section: benchmark

/*
// section: pprof-conversion
Total:       300ms      2.15s (flat, cum) 66.15%
     12            .          .           var global string // prevent benchmark inlining
     13            .          .
     14            .          .           func BenchmarkConversion(b *testing.B) {
     15            .          .           	var key = []byte{'J', 'o', 'h', 'n'}
     16            .          .           	var r string
     17         90ms       90ms           	for n := 0; n < b.N; n++ {
     18        210ms      2.06s           		r = beatles[string(key)]
     19            .          .           	}
     20            .          .           	global = r
     21            .          .           }
     22            .          .
     23            .          .           func BenchmarkAllocation(b *testing.B) {
// section: pprof-conversion

// section: pprof-allocation
  Total:       150ms      1.05s (flat, cum) 32.31%
     21            .          .           }
     22            .          .
     23            .          .           func BenchmarkAllocation(b *testing.B) {
     24            .          .           	var key = []byte{'J', 'o', 'h', 'n'}
     25            .          .           	var r string
     26         30ms       30ms           	for n := 0; n < b.N; n++ {
     27         40ms      310ms           		k := string(key)
     28         80ms      710ms           		r = beatles[k]
     29            .          .           	}
     30            .          .           	global = r
     31            .          .
     32            .          .           }
// section: pprof-allocation

// section: benchmark-output
BenchmarkConversion-8           200000000                8.46 ns/op            0 B/op          0 allocs/op
BenchmarkAllocation-8           100000000               12.4 ns/op             0 B/op          0 allocs/op
// section: benchmark-output

*/
