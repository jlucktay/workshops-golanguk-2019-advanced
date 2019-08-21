package main

import (
	"math/rand"
	"time"
)

// section: template
// Define a WaitGroup

func print(prefix string, count int) {
	// Defer the waitgroup Done

	for i := 0; i < count; i++ {
		randomSleep(50) // add some random delay to simulate better concurrency
		// Print out the prefix and the loop variable `i`
	}
}

// Define the wait group

func main() {
	// Increment WaitGroup

	// Call `print` with a Goroutine, first argument "first", second argument 50
	// Call `print` with a Goroutine, first argument "second", second argument 50

	// Wait...
}

// section: template

func randomSleep(i int) {
	// sleep for a random amount of time to add "chaos"
	time.Sleep(time.Duration(rand.Intn(i)) * time.Millisecond)
}
