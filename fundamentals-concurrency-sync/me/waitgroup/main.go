package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define a WaitGroup
var wg sync.WaitGroup

func print(prefix string, count int) {
	// Defer the waitgroup Done
	defer wg.Done()

	for i := 0; i < count; i++ {
		randomSleep(50) // add some random delay to simulate better concurrency
		// Print out the prefix and the loop variable `i`
		fmt.Printf("[%s] %d\n", prefix, i)
	}
}

func main() {
	// Increment WaitGroup
	wg.Add(2)

	// Call `print` with a Goroutine, first argument "first", second argument 50
	go print("first", 50)
	// Call `print` with a Goroutine, first argument "second", second argument 50
	go print("second", 50)

	// Wait...
	wg.Wait()
}

func randomSleep(i int) {
	// sleep for a random amount of time to add "chaos"
	time.Sleep(time.Duration(rand.Intn(i)) * time.Millisecond)
}
