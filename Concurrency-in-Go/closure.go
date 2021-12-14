package main

import (
	"fmt"
	"sync"
)

// You can think of a waitgroup like a concurrent safe counter. It calls increment the counter by the integer passed in,
// and calls to Done decrement the counter by one
// Calls to wait block until the counter is zero
// It's customary to couple calls to Add as closely as possible to the goroutines they're helping to track

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()

	//Correct format to print the contents of the string
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
