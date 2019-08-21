package benchmarking

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// section: code
func withPlus(x string) string {
	return x + x + x + x + x + x + x + x + x + x
}

func withSprintf(x string) string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", x, x, x, x, x, x, x, x, x, x)
}

func withBuffer(x string) string {
	bb := &bytes.Buffer{}
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	return bb.String()
}

func withStringBuilder(x string) string {
	bb := &strings.Builder{}
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	return bb.String()
}

// section: code

// section: bench
var sink string

func BenchmarkWithPlus(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = withPlus("hello")
	}
	sink = r
}

func BenchmarkWithSprintf(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = withSprintf("hello")
	}
	sink = r
}

func BenchmarkWithBuffer(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = withBuffer("hello")
	}
	sink = r
}

func BenchmarkWithBuilder(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = withStringBuilder("hello")
	}
	sink = r
}

// section: bench

/*
// section: output
goos: darwin
goarch: amd64
BenchmarkWithPlus-8             20000000                99.5 ns/op            64 B/op          1 allocs/op
BenchmarkWithSprintf-8           3000000               593 ns/op             224 B/op         11 allocs/op
BenchmarkWithBuffer-8           10000000               139 ns/op             128 B/op          2 allocs/op
BenchmarkWithBuilder-8          10000000               168 ns/op             120 B/op          4 allocs/op
PASS
ok      command-line-arguments  7.883s
// section: output
*/

// section: stretch
func BenchmarkAll(b *testing.B) {
	bms := []struct {
		name string
		f    func(string) string
	}{
		{
			name: "WithPlus",
			f:    withPlus,
		},
		{
			name: "WithSprintf",
			f:    withSprintf,
		},
		{
			name: "WithBuffer",
			f:    withBuffer,
		},
		{
			name: "WithStringBuilder",
			f:    withStringBuilder,
		},
	}

	for _, bm := range bms {
		b.Run(bm.name, func(b *testing.B) {
			var r string
			for n := 0; n < b.N; n++ {
				r = bm.f("hello")
			}
			sink = r
		})

	}
}

// section: stretch

/*
// section: stretch-output
goos: darwin
goarch: amd64
BenchmarkAll/WithPlus-8                 20000000               106 ns/op              64 B/op          1 allocs/op
BenchmarkAll/WithSprintf-8               3000000               596 ns/op             224 B/op         11 allocs/op
BenchmarkAll/WithBuffer-8               10000000               137 ns/op             128 B/op          2 allocs/op
BenchmarkAll/WithStringBuilder-8                10000000               170 ns/op             120 B/op          4 allocs/op
PASS
ok      command-line-arguments  8.033s
// section: stretch-output


*/
