package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func dowork(t time.Duration, res chan string, s string) {
	fmt.Println("\ndoing work...")
	time.Sleep(t)
	fmt.Println("work is done")

	res <- fmt.Sprintf("the result of the work %v", rand.Intn(100))

	wg.Done()
}

var wg *sync.WaitGroup

func main() {
	start := time.Now()
	res := make(chan string)

	wg = &sync.WaitGroup{}

	wg.Add(3)

	go dowork(time.Second*2, res, "work 1")
	go dowork(time.Second*4, res, "work 2")
	go dowork(time.Second*6, res, "work 3")

	go func() {

		for val := range res {
			fmt.Println(val)
		}
		fmt.Printf("Time took to get work done %v \n", time.Since(start))
	}()

	wg.Wait()
	close(res)
	time.Sleep(time.Second)
}
