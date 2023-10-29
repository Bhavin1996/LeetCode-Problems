package leetcode

import (
	"fmt"
	"sync"
)

func main_chan() {
	// Create a channel for passing values.
	ch := make(chan int)

	// Create a set of mutexes for synchronization.
	var mu sync.Mutex
	var m2 sync.Mutex
	var m3 sync.Mutex

	// Use a WaitGroup to ensure all goroutines complete before exiting.
	var wg sync.WaitGroup
	wg.Add(3)

	// Launch three goroutines.
	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			startValue := (id - 1) * 5
			endValue := id * 5
			for i := startValue + 1; i <= endValue; i++ {
				ch <- i
			}
		}(i)
	}

	// Close the channel once all goroutines are done.
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Print values from the channel in the desired order.
	for i := 1; i <= 15; i++ {
		select {
		case val := <-ch:
			switch val {
			case 1, 2, 3, 4, 5:
				mu.Lock()
				fmt.Println(val)
				mu.Unlock()
			case 6, 7, 8, 9, 10:
				m2.Lock()
				fmt.Println(val)
				m2.Unlock()
			case 11, 12, 13, 14, 15:
				m3.Lock()
				fmt.Println(val)
				m3.Unlock()
			}
		}
	}
}	var mu sync.Mutex

