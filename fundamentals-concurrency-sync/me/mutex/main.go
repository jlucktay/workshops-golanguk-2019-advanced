package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Numbers struct {
	// Add a sync.RWMutex
	rwm  sync.RWMutex
	data map[int]int
}

func (f *Numbers) Find(key int) (int, error) {
	// Protect this with a read lock
	f.rwm.RLock()
	defer f.rwm.RUnlock()
	k, ok := f.data[key]
	if !ok {
		return 0, fmt.Errorf("number not found: %d", key)
	}
	return k, nil
}

func (f *Numbers) Load() {
	var i int
	for {
		// Protect this with a write lock
		f.rwm.Lock()
		f.data[i] = i
		i++
		f.rwm.Unlock()
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

// randomSleep will sleep for a random amount of time (up to 'i' milliseconds) to add "chaos".
func randomSleep(i int) {
	time.Sleep(time.Duration(rand.Intn(i)) * time.Millisecond)
}
