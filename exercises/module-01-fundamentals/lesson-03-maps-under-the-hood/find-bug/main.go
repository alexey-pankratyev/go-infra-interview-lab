package main

import (
	"fmt"
	"sync"
)

func main() {
	counters := map[string]int{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("worker-%d", n%5)
			counters[key]++
		}(i)
	}

	wg.Wait()
	fmt.Println("done:", len(counters))
}
