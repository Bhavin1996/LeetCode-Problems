package leetcode

import "fmt"

func add(x chan int) {
	res := <-x + (<-x + 1)
	fmt.Println(res)
}

func main_add() {
	ch := make(chan int)
	for i := 0; i < 100; i++ {
		ch <- i
	}
	go add(ch)
}
