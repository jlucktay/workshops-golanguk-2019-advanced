package main

import (
	"fmt"
)

func main() {
	// section: main
	// Concatenate a string
	s := "hello "
	s = s + "world"

	// Append to a byte slice
	b := []byte("hello ")
	b = append(b, []byte("world")...)

	// print out both variables
	fmt.Printf("%s\n%s", s, string(b))

	// section: main
}

/*
// section: output
hello world
hello world
// section: output
*/
