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

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

func calcSquares(number int, squareop chan int) {
	//sum := 0
	//for number != 0 {
	//	digit := number % 10
	//	sum += digit * digit
	//	number /= 10
	//}
	//squareop <- sum

	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	//sum := 0
	//for number != 0 {
	//	digit := number % 10
	//	sum += digit * digit * digit
	//	number /= 10
	//}
	//cubeop <- sum

	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
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

	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares + cubes)

	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}

}