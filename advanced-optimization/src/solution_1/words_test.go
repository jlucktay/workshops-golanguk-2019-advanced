package words

import (
	"bytes"
	"io"
	"os"
	"testing"
)

var testdata = []byte(`
one
one two
one two three
one two three four
`)

func TestWords(t *testing.T) {
	tcs := []struct {
		name string
		f    func(io.Reader) (int, int, error)
	}{
		{
			name: "original",
			f:    Count,
		},
		{
			name: "bufio",
			f:    CountBufio,
		},
		{
			name: "bufio_bytes",
			f:    CountBufioBytes,
		},
		{
			name: "bufio_bytes_readbuff",
			f:    CountBufioBytesReadbuf,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := bytes.NewReader(testdata)
			c, d, err := tc.f(r)
			if err != nil {
				t.Fatal(err)
			}
			if c != 10 {
				t.Errorf("unexpected total count. got: %d, exp: %d", c, 10)
			}
			if d != 4 {
				t.Errorf("unexpected total distinct count. got: %d, exp: %d", d, 4)
			}
		})
	}
}

var sink int

func BenchmarkWords(b *testing.B) {
	var sum int

	bms := []struct {
		name string
		f    func(io.Reader) (int, int, error)
	}{
		{
			name: "original",
			f:    Count,
		},
		{
			name: "bufio",
			f:    CountBufio,
		},
		{
			name: "bufio_bytes",
			f:    CountBufioBytes,
		},
		{
			name: "bufio_bytes_readbuff",
			f:    CountBufioBytesReadbuf,
		},
	}
	for _, bm := range bms {
		b.Run(bm.name, func(b *testing.B) {
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				b.StopTimer()
				f, err := os.Open("./testdata/hamlet.txt")
				if err != nil {
					b.Fatal(err)
				}
				defer f.Close()

				b.StartTimer()

				c, d, err := bm.f(f)
				if err != nil {
					b.Fatal(err)
				}
				sum = c + d
			}
			sink = sum
		})
	}
}
