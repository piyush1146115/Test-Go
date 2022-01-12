package main

import (
	"fmt"
)

func main() {
	chanOwner := func() <-chan int {
		resultstream := make(chan int, 5)
		go func() {
			defer close(resultstream)
			for i := 0; i <= 5; i++ {
				resultstream <- i
			}
		}()
		return resultstream
	}

	resultstream := chanOwner()

	for result := range resultstream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
