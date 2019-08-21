package words

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

// section: count
// Count will return the total words and distinct words in a text file
func CountBufioBytes(rd io.Reader) (int, int, error) {
	// wrap the reader in a buffer
	br := bufio.NewReader(rd)

	var found bool
	words := map[string]int{}
	word := []rune{}
	count := 0

	for {
		r, err := readRune(br)
		if err == io.EOF {
			break
		}
		if err != nil {
			return -1, -1, fmt.Errorf("error reading: %v", err)
		}

		if unicode.IsSpace(r) && found {
			found = false
			words[string(word)] = words[string(word)] + 1
			word = word[:0] // clear outheslice
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
BenchmarkWords       100          11071109 ns/op         1425035 B/op     205645 allocs/op
// section: benchmark

// section: benchstat
name   old time/op    new time/op    delta
Words    12.2ms ± 1%    10.9ms ± 1%  -10.66%  (p=0.000 n=10+10)

name   old alloc/op   new alloc/op   delta
Words    1.89MB ± 0%    1.42MB ± 0%  -24.52%  (p=0.000 n=10+8)

name   old allocs/op  new allocs/op  delta
Words      290k ± 0%      206k ± 0%  -29.04%  (p=0.000 n=10+8)
// section: benchstat

*/
