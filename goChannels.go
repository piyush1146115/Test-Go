package main

import (
	"fmt"
	"time"
)

func hello1(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func main() {
	//a := make(chan int)
	//var a chan int
	//if a == nil {
	//	fmt.Println("channel a is nil, going to define it")
	//	a = make(chan int)
	//	fmt.Printf("Type of a is %T", a)
	//}

	done := make(chan bool)
	fmt.Println("Main going to call hello1 go goroutine")
	go hello1(done)
	//<-done
	fmt.Println("Main received data", <-done)
}