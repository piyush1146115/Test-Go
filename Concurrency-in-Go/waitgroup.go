package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()
	wg.Add(1) // Here we call Add with an argument of 1 to indicate that one goroutine is beginning.
	go func() {
		defer wg.Done() // Here we call Done using the defer keyword to ensure that before we exit the goroutine’s closure, we indicate to the WaitGroup that we’ve exited.
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()
	wg.Wait() //  Here we call Wait , which will block the main goroutine until all goroutines have indicated they have exited.
	fmt.Println("All goroutines complete.")

	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}
	const numGreeters = 5

	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()

	var count int
	increment := func() { count++ }
	decrement := func() { count-- }
	var once sync.Once
	once.Do(increment)
	once.Do(decrement)
	fmt.Printf("Count: %d\n", count)
}
