package words

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

// Count will return the total words and distinct words in a text file
func Count(rd io.Reader) (int, int, error) {
	var found bool
	words := map[string]int{}
	word := ""
	count := 0

	br := bufio.NewReader(rd)

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
