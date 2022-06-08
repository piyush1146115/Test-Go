package main

import "fmt"

func fn(m map[int]int) {
	m = make(map[int]int)
}

func main() {
	var m map[int]int
	fn(m)
	fmt.Println(m == nil)
}
