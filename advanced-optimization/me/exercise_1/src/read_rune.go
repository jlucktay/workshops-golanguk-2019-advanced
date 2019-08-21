package words

import (
	"io"
)

// section: readRune
func readRune(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

// section: readRune

// section: readRuneBuf
func readRuneBuf(r io.Reader, buf []byte) (rune, error) {
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

// section: readRuneBuf
