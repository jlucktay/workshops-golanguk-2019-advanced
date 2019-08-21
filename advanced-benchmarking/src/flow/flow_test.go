package flow_test

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// section: fatal
func BenchmarkFile(b *testing.B) {
	f, err := os.Open("testdata/nothing.txt")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	// continue on with benchmark...
	io.Copy(ioutil.Discard, f)
}

// section: fatal

// section: reset-timer
func BenchmarkFields(b *testing.B) {

	f, err := os.Open("testdata/hamlet.txt")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	// Reset benchmark timer
	// we don't want the overhead of opening a file benchmarked
	b.ResetTimer()

	// Run benchmark
	for i := 0; i < b.N; i++ {
		// continue on with benchmark...
		io.Copy(ioutil.Discard, f)
	}
}

// section: reset-timer

// section: stop-start-timer
func BenchmarkHamletStartStop(b *testing.B) {

	// Run benchmark
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		f, err := os.Open("testdata/hamlet.txt")
		if err != nil {
			b.Fatal(err)
		}
		defer f.Close()
		b.StartTimer()

		// continue on with benchmark...
		io.Copy(ioutil.Discard, f)
	}
}

// section: stop-start-timer

// section: sub
var sink int

func BenchmarkSub(b *testing.B) {
	// bms = benchmarks
	bms := []struct {
		name string
		data string
	}{
		{
			name: "one",
			data: "one",
		},
		{
			name: "two",
			data: "one two",
		},
		{
			name: "three",
			data: "one two three",
		},
	}

	for _, bm := range bms {
		b.Run(bm.name, func(b *testing.B) {
			var n int
			for i := 0; i < b.N; i++ {
				fields := strings.Fields(bm.data)
				n = len(fields)
			}
			sink = n
		})
	}
}

// section: sub

/*
// section: sub-output
BenchmarkSub/one-8              30000000                38.5 ns/op
BenchmarkSub/two-8              20000000                72.4 ns/op
BenchmarkSub/three-8            20000000                89.5 ns/op
// section: sub-output
*/
