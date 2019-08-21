// section: writer
package bufioex

import (
	"bufio"
	"io"
)

func NewWriter(w io.Writer) io.Writer {
	return &Writer{w: w}
}

type Writer struct {
	w       io.Writer
	writes  int
	written int
}

func (w *Writer) Write(p []byte) (int, error) {
	w.writes++
	w.written += len(p)
	return w.w.Write(p)
}

func (w *Writer) Writes() int  { return w.writes }
func (w *Writer) Written() int { return w.written }

// section: writer

func NewBufWriter(w io.Writer) io.Writer {
	w = &Writer{w: w}
	return bufio.NewWriter(w)
}

func NewBufWriterSize(w io.Writer, size int) io.Writer {
	w = &Writer{w: w}
	return bufio.NewWriterSize(w, size)
}
