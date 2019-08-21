package read

import (
	"bytes"
	"io"
	"testing"
)

var sink int

// section: read
func BenchmarkRead(b *testing.B) {
	data := make([]byte, 100000)
	for i := range data {
		data[i] = byte(i)
	}

	b.ResetTimer()
	var count int
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		r := bytes.NewReader(data)
		b.StartTimer()

		for {
			bt, err := readBytes(r)
			if err == io.EOF {
				break
			}
			if err != nil {
				b.Fatal(err)
			}
			count = len(bt)
		}
		sink = count
	}
}

// section: read

// section: read-buf
func BenchmarkReadBuf(b *testing.B) {
	data := make([]byte, 100000)
	for i := range data {
		data[i] = byte(i)
	}

	b.ResetTimer()
	var count int
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		r := bytes.NewReader(data)
		buf := make([]byte, 8)
		b.StartTimer()

		for {
			bt, err := readBytesBuf(buf, r)
			if err == io.EOF {
				break
			}
			if err != nil {
				b.Fatal(err)
			}
			count = len(bt)
		}
		sink = count
	}
}

// section: read-buf

/*
// section: read-output
BenchmarkRead-8             5000            274444 ns/op          100016 B/op      12501 allocs/op

// section: read-output

// section: read-buf-output
BenchmarkReadBuf-8         20000             89190 ns/op               0 B/op          0 allocs/op
// section: read-buf-output

// section: benchstat
name    old time/op    new time/op    delta
Read-8     276µs ± 1%      94µs ± 7%   -65.83%  (p=0.000 n=9+10)

name    old alloc/op   new alloc/op   delta
Read-8     100kB ± 0%       0kB       -100.00%  (p=0.000 n=10+10)

name    old allocs/op  new allocs/op  delta
Read-8     12.5k ± 0%      0.0k       -100.00%  (p=0.000 n=10+10)
// section: benchstat

*/
