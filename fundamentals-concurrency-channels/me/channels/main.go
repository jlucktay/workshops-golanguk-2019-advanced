package main

import "fmt"

func process(messages chan string, quit chan struct{}) {
	// Using the `range` operator, loop through messages
	// until the channel is closed
	for m := range messages {
		// print the message with a for loop using range
		fmt.Printf("message: %s\n", m)
	}

	// close the `quit` channel
	close(quit)
}

func main() {
	// declare the messages channel of type string and capacity of 5
	messages := make(chan string, 5)

	// declare a signal channel
	quit := make(chan struct{})

	// launch the `process` function in a goroutine
	go process(messages, quit)

	// declare 5 fruits in a []string
	fruits := []string{"apple", "banana", "watermelon", "strawberry", "pear"}

	// loop through the fruits and send them to the messages channel
	for _, f := range fruits {
		messages <- f
	}

	// close the messages channel
	close(messages)

	// wait for everything to finish (hint, block on the quit channel)
	<-quit

	fmt.Println("done")
}
