package read

import "io"

// section: read
func readBytes(r io.Reader) ([]byte, error) {
	buf := make([]byte, 8)
	_, err := r.Read(buf)
	return buf, err
}

// section: read

// section: read-buf
func readBytesBuf(buf []byte, r io.Reader) ([]byte, error) {
	_, err := r.Read(buf)
	return buf, err
}

// section: read-buf
