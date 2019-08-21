package words

import (
	"bytes"
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
	r := bytes.NewReader(testdata)
	c, d, err := Count(r)
	if err != nil {
		t.Fatal(err)
	}
	if c != 10 {
		t.Errorf("unexpected total count. got: %d, exp: %d", c, 10)
	}
	if d != 4 {
		t.Errorf("unexpected total distinct count. got: %d, exp: %d", d, 4)
	}
}

var sink int

// section: benchmark
func BenchmarkWords(b *testing.B) {
	var sum int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		f, err := os.Open("./testdata/hamlet.txt")
		if err != nil {
			b.Fatal(err)
		}
		defer f.Close()

		b.StartTimer()

		c, d, err := Count(f)
		if err != nil {
			b.Fatal(err)
		}
		sum = c + d
	}
	sink = sum
}

// section: benchmark
