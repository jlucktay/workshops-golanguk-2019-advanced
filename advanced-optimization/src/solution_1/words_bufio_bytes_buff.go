package words

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

// section: count
// Count will return the total words and distinct words in a text file
func CountBufioBytesReadbuf(rd io.Reader) (int, int, error) {
	// wrap the reader in a buffer
	br := bufio.NewReader(rd)

	var found bool
	words := map[string]int{}
	count := 0
	word := []rune{}

	buf := make([]byte, 1)
	for {
		r, err := readRuneBuf(br, buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return -1, -1, fmt.Errorf("error reading: %v", err)
		}

		if unicode.IsSpace(r) && found {
			found = false
			c := words[string(word)]
			c++
			words[string(word)] = c
			word = word[:0]
			count++
		}
		found = unicode.IsLetter(r)
		if found {
			word = append(word, r)
		}
	}
	return count, len(words), nil
}

// section: count

/*
// section: benchmark
BenchmarkWords       200           9023395 ns/op         1244294 B/op      25419 allocs/op
// section: benchmark

// section: benchstat
name   old time/op    new time/op    delta
Words    10.9ms ± 1%     8.8ms ± 1%  -19.01%  (p=0.000 n=10+10)

name   old alloc/op   new alloc/op   delta
Words    1.42MB ± 0%    1.24MB ± 0%  -12.64%  (p=0.000 n=8+10)

name   old allocs/op  new allocs/op  delta
Words      206k ± 0%       25k ± 0%  -87.64%  (p=0.000 n=8+10)

// section: benchstat


// section: benchmark-final
BenchmarkWords         3         499446255 ns/op         1892213 B/op     289795 allocs/op
BenchmarkWords       200           9023395 ns/op         1244294 B/op      25419 allocs/op
// section: benchmark-final

// section: benchstat-final

name   old time/op    new time/op    delta
Words     512ms ± 3%       9ms ± 1%  -98.27%  (p=0.000 n=9+10)

name   old alloc/op   new alloc/op   delta
Words    1.89MB ± 0%    1.24MB ± 0%  -34.16%  (p=0.000 n=10+10)

name   old allocs/op  new allocs/op  delta
Words      290k ± 0%       25k ± 0%  -91.23%  (p=0.000 n=10+10)

// section: benchstat-final

*/
