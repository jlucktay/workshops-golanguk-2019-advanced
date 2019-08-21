package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// section: solution
var wg sync.WaitGroup

func print(prefix string, count int) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		randomSleep(50) // add some random delay to simulate better concurrency
		fmt.Println(prefix, i)
	}
}

func main() {
	wg.Add(2)
	go print("first: ", 50)
	go print("second: ", 50)
	wg.Wait()
}

// section: solution

func randomSleep(i int) {
	// sleep for a random amount of time to add "chaos"
	time.Sleep(time.Duration(rand.Intn(i)) * time.Millisecond)
}
