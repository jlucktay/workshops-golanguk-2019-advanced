package main

import "fmt"

func main() {
	c := make(<-chan string)
	go func() {
		c <- "foo"
	}()
	fmt.Println(<-c)
}
