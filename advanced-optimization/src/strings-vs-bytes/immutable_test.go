package main

import "testing"

// section: benchmark
var str string
var bytes []byte

func BenchmarkString(b *testing.B) {
	cat := func(s1, s2 string) string {
		return s1 + s2
	}
	hello := "hello "
	world := "world"

	var s1 string
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s := cat(hello, world)
		s1 = s
	}
	str = s1
}

func BenchmarkBytes(b *testing.B) {
	bcat := func(b1, b2 []byte) []byte {
		return append(b1, b2...)
	}

	hello := []byte("hello ")
	world := []byte("world")

	var b1 []byte
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bt := bcat(hello, world)
		b1 = bt
	}
	bytes = b1
}

// section: benchmark

/*

// section: command
go test -benchmem -bench=. ./immutable_test.go
// section: command

// section: output
BenchmarkString-8       30000000                40.8 ns/op            16 B/op          1 allocs/op
BenchmarkBytes-8        50000000                28.1 ns/op            16 B/op          1 allocs/op
// section: output

// section: pprof commands
go test -cpuprofile=cpu-string.pprof -bench=String ./immutable_test.go
go test -cpuprofile=cpu-bytes.pprof -bench=Bytes ./immutable_test.go
go test -memprofile=mem-string.pprof -bench=String ./immutable_test.go
go test -memprofile=mem-bytes.pprof -bench=Bytes ./immutable_test.go

pprof --alloc_objects -http=:8181 mem.pprof
*/

/*

// section: pprof-bytes
command-line-arguments.BenchmarkBytes.func1

  Total:    776.01MB   776.01MB (flat, cum) 62.56%
     22            .          .           	str = s1
     23            .          .           }
     24            .          .
     25            .          .           func BenchmarkBytes(b *testing.B) {
     26            .          .           	bcat := func(b1, b2 []byte) []byte {
     27     776.01MB   776.01MB           		return append(b1, b2...)
     28            .          .           	}
     29            .          .
     30            .          .           	hello := []byte("hello ")
     31            .          .           	world := []byte("world")
     32            .          .
// section: pprof-bytes

// section: pprof-strings
command-line-arguments.BenchmarkString.func1

  Total:    464.51MB   464.51MB (flat, cum) 37.44%
      6            .          .           var str string
      7            .          .           var bytes []byte
      8            .          .
      9            .          .           func BenchmarkString(b *testing.B) {
     10            .          .           	cat := func(s1, s2 string) string {
     11     464.51MB   464.51MB           		return s1 + s2
     12            .          .           	}
     13            .          .           	hello := "hello "
     14            .          .           	world := "world"
     15            .          .
     16            .          .           	var s1 string
// section: pprof-strings
*/
