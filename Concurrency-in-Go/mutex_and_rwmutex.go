package main

import (
	"fmt"
	"sync"
)

func main() {
	// A Mutex provides a concurrent-safe way to express exclusive access to these shared resources

	// Increment

	//arithmetic.Wait()
	//fmt.Println("Arithmetic complete.")

	//Here we request exclusive use of the critical section - in this case the count variable - guarded by a Mutex, lock

	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	var arithmetic sync.WaitGroup

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")

	//The sync.RQMutex is conceptually the same thing as a Mutex: it guards access to memory
	// however, RWMutex gives you a little bit more control over the memory.
	// You can request a lock for reading, in which case you will be granted access unless the lock is being held for writing
	// This means an arbitrary number of readers can hold a reader lock so long as nothing else is holding a writer lock

	//It's usually advisable to  use RWMutex instead of Mutex when it logically makes sense

}
