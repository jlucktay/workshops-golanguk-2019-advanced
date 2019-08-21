package main

import (
	"fmt"
	"os"

	bufioex "github.com/gopherguides/learn/_training/advanced/optimization/src/bufio_ex"
)

func main() {
	// section: main
	w := bufioex.NewWriter(os.Stdout)

	// write the data
	data := []byte("abcdefghijklmnopqrstuvwxyz\n")
	for _, v := range data {
		w.Write([]byte{v})
	}
	fmt.Printf("Writes: %d, Bytes Written: %d", w.(*bufioex.Writer).Writes(), w.(*bufioex.Writer).Written())
	// section: main
}

/*
// section: output
abcdefghijklmnopqrstvwxyz
Writes: 27, Bytes Written: 27
// section: output
*/
