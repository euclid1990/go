package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string, gr int) {
	// Lock so only one goroutine at a time can access the map c.v.
	c.mux.Lock()
	c.v[key]++
	// fmt.Println("Go routine: ", gr)
	// Unlock for an other can access the map c.v.
	defer c.mux.Unlock()
}

// Value returns the current value of the counter for given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100000; i++ {
		go c.Inc("somekey", i)
		if i%10000 == 0 {
			fmt.Println(c.Value("somekey"))
		}
	}
	time.Sleep(2 * time.Second)
}
