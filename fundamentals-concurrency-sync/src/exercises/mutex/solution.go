package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// section: solution
type Numbers struct {
	sync.RWMutex
	data map[int]int
}

func (f *Numbers) Find(key int) (int, error) {
	f.RLock()
	defer f.RUnlock()
	k, ok := f.data[key]
	if !ok {
		return 0, fmt.Errorf("number not found: %d", key)
	}
	return k, nil
}

func (f *Numbers) Load() {
	var i int
	for {
		f.Lock()
		f.data[i] = i
		f.Unlock()
		i++
		randomSleep(75) // simulate real world load
	}
}

func NewNumbers() *Numbers {
	f := &Numbers{
		data: map[int]int{},
	}
	go f.Load()
	return f
}

var wg sync.WaitGroup

func main() {
	f := NewNumbers()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			for {
				_, err := f.Find(i)
				if err == nil {
					fmt.Printf("found %d\n", i)
					break
				}
				fmt.Println(err)
				randomSleep(25) // create some delay in our goroutine
			}
		}
	}()
	wg.Wait()
}

// section: solution

func randomSleep(i int) {
	// sleep for a random amount of time to add "chaos"
	time.Sleep(time.Duration(rand.Intn(i)) * time.Millisecond)
}
