package main

import (
	"bufio"
	"fmt"
	"os"

	bufioex "github.com/gopherguides/learn/_training/advanced/optimization/src/bufio_ex"
)

func main() {
	// section: main
	w := bufioex.NewWriter(os.Stdout)

	// Wrap the writer in a bufio.Writer
	bw := bufio.NewWriter(w)

	// write the data
	data := []byte("abcdefghijklmnopqrstuvwxyz\n")
	for _, v := range data {
		bw.Write([]byte{v})
	}
	// be sure to call flush to empty the buffer
	bw.Flush()

	fmt.Printf(
		"Writes: %d, Bytes Written: %d",
		w.(*bufioex.Writer).Writes(),  // don't do blind assertions in production!
		w.(*bufioex.Writer).Written(), // this is a benchmark so cheating is ok...
	)
	// section: main
}

/*
// section: output
abcdefghijklmnopqrstvwxyz
Writes: 1, Bytes Written: 27
// section: output
*/
