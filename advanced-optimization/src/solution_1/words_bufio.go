package words

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

// section: count
// Count will return the total words and distinct words in a text file
func CountBufio(rd io.Reader) (int, int, error) {
	// section: bufio
	// wrap the reader in a buffer
	br := bufio.NewReader(rd)
	// section: bufio

	var found bool
	words := map[string]int{}
	word := ""
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
			words[word] = words[word] + 1
			word = ""
			count++
		}
		found = unicode.IsLetter(r)
		if found {
			word += string(r)
		}
	}
	return count, len(words), nil
}

// section: count

/*
// section: benchmark
BenchmarkWords       100          12498754 ns/op         1886970 B/op     289780 allocs/op
// section: benchmark

// section: benchstat
name   old time/op    new time/op    delta
Words     512ms ± 3%      12ms ± 1%  -97.61%  (p=0.000 n=9+10)

name   old alloc/op   new alloc/op   delta
Words    1.89MB ± 0%    1.89MB ± 0%   -0.15%  (p=0.000 n=10+10)

name   old allocs/op  new allocs/op  delta
Words      290k ± 0%      290k ± 0%     ~     (p=0.189 n=10+10)
// section: benchstat

*/
