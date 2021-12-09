package main

import (
	"fmt"
	"sync"
)

//https://go.dev/blog/maps

func main() {
	var m map[string]int
	m = make(map[string]int)
	m["route"] = 66

	//This statement retrieves the value stored under the key "route" and assigns it to a new variable i:
	i := m["route"]
	fmt.Println(i)

	j := m["root"]
	fmt.Printf("%v", j)

	//In this statement, the first value (i) is assigned the value stored under the key "route".
	//If that key doesn’t exist, i is the value type’s zero value (0). The second value (ok) is a bool that is true if the key exists in the map, and false if not.
	i, ok := m["route"]
	fmt.Println(ok)

	//To initialize a map with some data, use a map literal:
	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}

	for key, value := range commits {
		fmt.Println("Key:", key, "Value:", value)
	}

	//Maps are not safe for concurrent use: it’s not defined what happens when you read and write to them simultaneously.
	//One common way to protect maps is with sync.RWMutex.
	//This statement declares a counter variable that is an anonymous struct containing a map and an embedded sync.RWMutex.

	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	//To write to the counter, take the write lock:
	counter.Lock()
	counter.m["some_key"]++
	counter.Unlock()

	//To read from the counter, take the read lock:
	counter.RLock()
	n := counter.m["some_key"]
	counter.RUnlock()
	fmt.Println("some_key:", n)

}
