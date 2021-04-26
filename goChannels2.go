package main

//func sendData(sendch chan<- int) {
//	sendch <- 10
//}
//
//func main() {
//	chnl := make(chan int)
//	go sendData(chnl)
//	fmt.Println(<-chnl)
//}

//If ok is false it means that we are reading from a closed channel. The value read from a closed channel will be the zero value of the channel's type.
//For example, if the channel is an int channel, then the value received from a closed channel will be 0.

//func produce(chnl chan int) {
//	for i := 0; i < 10; i++ {
//		chnl <- i
//	}
//	close(chnl)
//}
//
//func main() {
//	ch := make(chan int)
//	go produce(ch)
//	for {
//		v, ok := <-ch
//		if ok == false {
//			break
//		}
//		fmt.Println("Received ", v, ok)
//	}
//}

//Once a channel has been closed, you cannot send a value on this channel, but you can still receive from the channel.

import "fmt"

func main() {
	ch := make(chan bool, 2)
	ch <- true
	ch <- true
	close(ch)

	for i := 0; i < cap(ch)+1; i++ {
		v, ok := <-ch
		fmt.Println(v, ok)
	}
}
