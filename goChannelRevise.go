package main
//
//import "fmt"
//
//func main() {
//	ch := make(chan bool, 2)
//	ch <- true
//	ch <- true
//	close(ch)
//
//	//we retrieve the first two values we sent on the channel,
//	//then on our third attempt the channel gives us the values of false and false.
//	//The first false is the zero value for that channel’s type, which is false, as the channel is of type chan bool.
//	//The second indicates the open state of the channel, which is now false
//	for i := 0; i < cap(ch) +1 ; i++ {
//		v, ok := <- ch
//		fmt.Println(v, ok)
//	}
//
//	//for v := range ch {
//	//	fmt.Println(v) // called twice
//	//}
//
//
//	//nil channels blocks forever
//	var ch2 chan bool
//	<- ch2 // blocks forever
//}


import (
	"fmt"
	"time"
)

var cnt int
func WaitMany(a, b chan bool) {
	for a != nil || b != nil {
		select {
		case <-a:
			a = nil
		case <-b:
			b = nil
		}
		cnt++
		fmt.Println("From WaitMany ", cnt)
	}
}

func main() {
	a, b := make(chan bool), make(chan bool)
	t0 := time.Now()
	cnt = 0
	go func() {
		close(a)
		close(b)
	}()
	WaitMany(a, b)
	fmt.Printf("waited %v for WaitMany\n", time.Since(t0))
}