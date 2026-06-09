package main

import (
	"fmt"
	"sync"
)

type CounterMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func NewCounterMap() *CounterMap {
	return &CounterMap{m: make(map[string]int)}
}

// Inc increments the counter for key by 1.
func (c *CounterMap) Inc(key string) {
	// TODO: implement me.
}

// Get returns the current count for key. Returns 0 if key does not exist.
func (c *CounterMap) Get(key string) int {
	// TODO: implement me.
	return 0
}

// Top returns the key with the highest count and its count.
// Returns ("", 0) if the map is empty.
func (c *CounterMap) Top() (string, int) {
	// TODO: implement me.
	return "", 0
}

func main() {
	cm := NewCounterMap()
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("worker-%d", n%5)
			cm.Inc(key)
		}(i)
	}

	wg.Wait()

	topKey, topCount := cm.Top()
	fmt.Printf("top key: %s  count: %d\n", topKey, topCount)

	total := 0
	for i := 0; i < 5; i++ {
		total += cm.Get(fmt.Sprintf("worker-%d", i))
	}
	fmt.Println("total:", total)
}
