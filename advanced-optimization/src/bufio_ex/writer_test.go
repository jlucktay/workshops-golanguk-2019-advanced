package bufioex_test

import (
	"bufio"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	bufioex "github.com/gopherguides/learn/_training/advanced/optimization/src/bufio_ex"
)

func BenchmarkUnbuffered(b *testing.B) {
	w := bufioex.NewWriter(ioutil.Discard)
	data := []byte("abcdefghijklmnopqrstuvwxyz\n")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, v := range data {
			w.Write([]byte{v})
		}
	}
}

func BenchmarkBuffered(b *testing.B) {
	w := bufioex.NewBufWriter(ioutil.Discard)
	data := []byte("abcdefghijklmnopqrstuvwxyz\n")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, v := range data {
			w.Write([]byte{v})
		}
	}
	w.(*bufio.Writer).Flush()
}

func BenchmarkBufferedSize(b *testing.B) {
	w := bufioex.NewBufWriterSize(ioutil.Discard, 8)
	data := []byte("abcdefghijklmnopqrstuvwxyz\n")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, v := range data {
			w.Write([]byte{v})
		}
	}
	w.(*bufio.Writer).Flush()
}

// Write some buffered tests

func loadHamlet() (io.ReadCloser, error) {
	r, err := http.Get("http://hamlet.gopherguides.com")
	if err != nil {
		return nil, err
	}
	return r.Body, nil
}

// section: hamlet-unbuffered
func BenchmarkHamletUnbuffered(b *testing.B) {
	r, err := loadHamlet()
	defer r.Close()
	if err != nil {
		b.Fatal(err)
	}
	f, err := os.Create("./testdata/hamlet-unbuffered.txt")
	if err != nil {
		b.Fatal(err)
	}
	w := bufioex.NewWriter(f)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if _, err := io.Copy(w, r); err != nil {
			b.Fatal(err)
		}
	}
	//b.Logf("Writes: %d, Bytes Written: %d", w.(*bufioex.Writer).Writes(), w.(*bufioex.Writer).Written())
}

// section: hamlet-unbuffered

// section: hamlet-buffered
func BenchmarkHamletBuffered(b *testing.B) {
	r, err := loadHamlet()
	defer r.Close()
	if err != nil {
		b.Fatal(err)
	}

	f, err := os.Create("./testdata/hamlet-buffered.txt")
	if err != nil {
		b.Fatal(err)
	}
	w := bufioex.NewWriter(f)

	// Wrap the writer in a bufio.Writer
	bw := bufio.NewWriter(w)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if _, err := io.Copy(bw, r); err != nil {
			b.Fatal(err)
		}
	}
	//b.Logf("Writes: %d, Bytes Written: %d", w.(*bufioex.Writer).Writes(), w.(*bufioex.Writer).Written())
}

// section: hamlet-buffered

/*

run command:
$ go test -benchmem  -bench=BenchmarkHamlet

// section: results

BenchmarkHamletUnbuffered-8    500000  3171.0 ns/op  32768 B/op  1 allocs/op
BenchmarkHamletBuffered-8    20000000    64.1 ns/op      0 B/op  0 allocs/op
// section: results

// section: writes
--- BENCH: BenchmarkHamletUnbuffered-8
    writer_test.go:71: Writes: 85, Bytes Written: 182399
    writer_test.go:71: Writes: 90, Bytes Written: 182399
    writer_test.go:71: Writes: 88, Bytes Written: 182399
    writer_test.go:71: Writes: 87, Bytes Written: 182399
    writer_test.go:71: Writes: 91, Bytes Written: 182399
    writer_test.go:71: Writes: 84, Bytes Written: 182399
    writer_test.go:71: Writes: 89, Bytes Written: 182399

--- BENCH: BenchmarkHamletBuffered-8
    writer_test.go:94: Writes: 44, Bytes Written: 180224
    writer_test.go:94: Writes: 44, Bytes Written: 180224
    writer_test.go:94: Writes: 44, Bytes Written: 180224
    writer_test.go:94: Writes: 44, Bytes Written: 180224
    writer_test.go:94: Writes: 44, Bytes Written: 180224
    writer_test.go:94: Writes: 44, Bytes Written: 180224
    writer_test.go:94: Writes: 44, Bytes Written: 180224

// section: writes
*/
