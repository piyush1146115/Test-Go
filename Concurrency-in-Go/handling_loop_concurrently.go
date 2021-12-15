package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	sum := make(chan int)

	for n := range make([]int, 1000) {
		wg.Add(1)
		go func(n int, ch chan int) {
			ch <- n
			wg.Done()
		}(n, ch)
	}

	go Add(ch, sum)

	wg.Wait()
	close(ch)
	fmt.Println(<-sum)
	close(sum)
}

func Add(c, sum chan int) {
	s := 0
	for {
		v, ok := <-c
		if !ok {
			break
		}
		s += v
	}
	sum <- s
}
